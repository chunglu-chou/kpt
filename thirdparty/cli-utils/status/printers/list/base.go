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
	"strings"

	"sigs.k8s.io/cli-utils/pkg/apply/event"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/collector"
	pollevent "sigs.k8s.io/cli-utils/pkg/kstatus/polling/event"
	"sigs.k8s.io/cli-utils/pkg/object"
	"sigs.k8s.io/cli-utils/pkg/print/common"
	"sigs.k8s.io/cli-utils/pkg/print/list"
)

// BaseListPrinter implements the Printer interface and outputs the resource
// status information as a list of events as they happen.
type BaseListPrinter struct {
	Formatter list.Formatter
	Data      *PrintData
}

// PrintData records data required for printing
type PrintData struct {
	IndexResourceMap map[int]object.ObjMetadata
	IndexGroupMap    map[int]string
	MaxElement       int
	Identifiers      object.ObjMetadataSet
	StatusSet        map[string]bool
}

func (ep *BaseListPrinter) PrintStatus(coll *collector.ResourceStatusCollector, id object.ObjMetadata) error {
	for idx := 0; idx < ep.Data.MaxElement; idx++ {
		if text, ok := ep.Data.IndexGroupMap[idx]; ok {
			// this index represents header for each inventory name
			fmt.Println(text)
		} else {
			// this index represents an object on the cluster
			identifier := ep.Data.IndexResourceMap[idx]
			// retrieve the status of object
			status := coll.ResourceStatuses[identifier]
			// check if the status is filtered out
			if _, ok := ep.Data.StatusSet[strings.ToLower(status.Status.String())]; !ok && len(ep.Data.StatusSet) != 0 {
				continue
			}
			// if the object is the one triggered the event channel, make the text green
			if identifier == id {
				fmt.Printf("%c[%dm\r", common.ESC, common.GREEN)
			}
			// print out status
			err := ep.Formatter.FormatStatusEvent(event.StatusEvent{
				Identifier:       identifier,
				PollResourceInfo: status,
				Resource:         status.Resource,
			})
			// reset the color of text if set to other colors
			if identifier == id {
				fmt.Printf("%c[%dm\r", common.ESC, common.RESET)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// PrintError print out errors when received error events
func (ep *BaseListPrinter) PrintError(e error) error {
	err := ep.Formatter.FormatErrorEvent(event.ErrorEvent{Err: e})
	if err != nil {
		return err
	}
	return nil
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
			// move the cursor to the origin
			fmt.Printf("%c[H", common.ESC)
			// clear all printed content
			fmt.Printf("%c[0J\r", common.ESC)
			err := ep.PrintStatus(coll, e.Resource.Identifier)
			if err != nil {
				panic(err)
			}
			if e.Type == pollevent.ErrorEvent {
				err := ep.PrintError(e.Error)
				if err != nil {
					panic(err)
				}
			}
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
