// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

import (
	"fmt"
	"sync"

	"sigs.k8s.io/cli-utils/pkg/apply/event"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/collector"
	pollevent "sigs.k8s.io/cli-utils/pkg/kstatus/polling/event"
	"sigs.k8s.io/cli-utils/pkg/object"
	"sigs.k8s.io/cli-utils/pkg/print/list"
)

const (
	// printing parameters
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorReset  = "\033[0m"
	separator   = "-----------------------------------------"
)

// BaseListPrinter implements the Printer interface and outputs the resource
// status information as a list of events as they happen.
type BaseListPrinter struct {
	Formatter list.Formatter
	InvName   string
	Lock      sync.Mutex
}

// Print takes an event channel and outputs the status events on the channel
// until the channel is closed. The provided cancelFunc is consulted on
// every event and is responsible for stopping the poller when appropriate.
// This function will block.
func (ep *BaseListPrinter) Print(ch <-chan pollevent.Event, identifiers []object.ObjMetadata,
	cancelFunc collector.ObserverFunc) error {
	coll := collector.NewResourceStatusCollector(identifiers)
	// The actual work is done by the collector, which will invoke the
	// callback on every event. In the callback we print the status
	// information and call the cancelFunc which is responsible for
	// stopping the poller at the correct time.
	done := coll.ListenWithObserver(ch, collector.ObserverFunc(
		func(statusCollector *collector.ResourceStatusCollector, e pollevent.Event) {
			// critical section
			ep.Lock.Lock()
			// print the inventory name
			fmt.Println(separator)
			fmt.Println(colorCyan + ep.InvName + colorReset)
			fmt.Println(separator)
			fmt.Print(colorYellow)
			err := ep.printStatusEvent(e)
			fmt.Print(colorReset)
			if err != nil {
				// error detected, quit critical section
				ep.Lock.Unlock()
				panic(err)
			}
			// print all status of resources except the latest updated one under this inventory group
			for id, status := range statusCollector.ResourceStatuses {
				if id != e.Resource.Identifier {
					err := ep.Formatter.FormatStatusEvent(event.StatusEvent{
						Identifier:       id,
						Resource:         status.Resource,
						PollResourceInfo: status,
					})
					if err != nil {
						// error detected, quit critical section
						ep.Lock.Unlock()
						panic(err)
					}
				}
			}
			// printing ended, quit critical section
			ep.Lock.Unlock()
			cancelFunc(statusCollector, e)
		}),
	)
	// Block until the done channel is closed.
	<-done
	if o := coll.LatestObservation(); o.Error != nil {
		return o.Error
	}
	return nil
}

func (ep *BaseListPrinter) printStatusEvent(se pollevent.Event) error {
	switch se.Type {
	case pollevent.ResourceUpdateEvent:
		id := se.Resource.Identifier
		return ep.Formatter.FormatStatusEvent(event.StatusEvent{
			Identifier:       id,
			Resource:         se.Resource.Resource,
			PollResourceInfo: se.Resource,
		})
	case pollevent.ErrorEvent:
		return ep.Formatter.FormatErrorEvent(event.ErrorEvent{
			Err: se.Error,
		})
	}
	return nil
}
