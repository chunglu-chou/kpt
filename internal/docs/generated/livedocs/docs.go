// Code generated by "mdtogo"; DO NOT EDIT.
package livedocs

var LiveShort = `Deploy local packages to a cluster.`
var LiveLong = `
The ` + "`" + `live` + "`" + ` command group contains subcommands for deploying local
` + "`" + `kpt` + "`" + ` packages to a cluster.
`

var ApplyShort = `Apply a package to the cluster (create, update, prune).`
var ApplyLong = `
  kpt live apply [PKG_PATH | -] [flags]

Args:

  PKG_PATH | -:
    Path to the local package which should be applied to the cluster. It must
    contain a Kptfile or a ResourceGroup manifest with inventory metadata.
    Defaults to the current working directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:

  --dry-run:
    It true, kpt will validate the resources in the package and print which
    resources will be applied and which resources will be pruned, but no resources
    will be changed.
    If the --server-side flag is true, kpt will do a server-side dry-run, otherwise
    it will be a client-side dry-run. Note that the output will differ somewhat
    between the two alternatives.
  
  --field-manager:
    Identifier for the **owner** of the fields being applied. Only usable
    when --server-side flag is specified. Default value is kubectl.
  
  --force-conflicts:
    Force overwrite of field conflicts during apply due to different field
    managers. Only usable when --server-side flag is specified.
    Default value is false (error and failure when field managers conflict).
  
  --install-resource-group:
    Install the ResourceGroup CRD into the cluster if it isn't already
    available. Default is false.
  
  --inventory-policy:
    Determines how to handle overlaps between the package being currently applied
    and existing resources in the cluster. The available options are:
  
      * strict: If any of the resources already exist in the cluster, but doesn't
        belong to the current package, it is considered an error.
      * adopt: If a resource already exist in the cluster, but belongs to a
        different package, it is considered an error. Resources that doesn't belong
        to other packages are adopted into the current package.
  
    The default value is ` + "`" + `strict` + "`" + `.
  
  --output:
    Determines the output format for the status information. Must be one of the following:
  
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
  
  --poll-period:
    The frequency with which the cluster will be polled to determine
    the status of the applied resources. The default value is 2 seconds.
  
  --prune-propagation-policy:
    The propagation policy that should be used when pruning resources. The
    default value here is 'Background'. The other options are 'Foreground' and 'Orphan'.
  
  --prune-timeout:
    The threshold for how long to wait for all pruned resources to be
    deleted before giving up. If this flag is not set, kpt live apply will wait
    until interrupted. In most cases, it would also make sense to set the
    --prune-propagation-policy to Foreground when this flag is set.
  
  --reconcile-timeout:
    The threshold for how long to wait for all resources to reconcile before
    giving up. If this flag is not set, kpt live apply will wait until
    interrupted.
  
  --server-side:
    Perform the apply operation server-side rather than client-side.
    Default value is false (client-side).
  
  --show-status-events:
    The output will include the details on the reconciliation status
    for all resources. Default is ` + "`" + `false` + "`" + `.
  
    Does not apply for the ` + "`" + `table` + "`" + ` output format.
`
var ApplyExamples = `
  # apply resources in the current directory
  $ kpt live apply

  # apply resources in the my-dir directory and wait up until 15 minutes 
  # for all the resources to be reconciled before pruning
  $ kpt live apply --reconcile-timeout=15m my-dir

  # apply resources and specify how often to poll the cluster for resource status
  $ kpt live apply --reconcile-timeout=15m --poll-period=5s my-dir
`

var DestroyShort = `Remove all previously applied resources in a package from the cluster`
var DestroyLong = `
  kpt live destroy [PKG_PATH | -]

Args:

  PKG_PATH | -:
    Path to the local package which should be deleted from the cluster. It must
    contain a Kptfile or a ResourceGroup manifest with inventory metadata.
    Defaults to the current working directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:

  --dry-run:
    It true, kpt will print the resources that will be removed from the cluster,
    but no resources will be deleted.
  
  --inventory-policy:
    Determines how to handle overlaps between the package being currently applied
    and existing resources in the cluster. The available options are:
  
      * strict: If any of the resources already exist in the cluster, but doesn't
        belong to the current package, it is considered an error.
      * adopt: If a resource already exist in the cluster, but belongs to a
        different package, it is considered an error. Resources that doesn't belong
        to other packages are adopted into the current package.
  
    The default value is ` + "`" + `strict` + "`" + `.
  
  --output:
    Determines the output format for the status information. Must be one of the following:
  
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
  
  --show-status-events:
    The output will include the details on the reconciliation status
    for all resources. Default is ` + "`" + `false` + "`" + `.
  
    Does not apply for the ` + "`" + `table` + "`" + ` output format.
`
var DestroyExamples = `
  # remove all resources in the current package from the cluster.
  $ kpt live destroy
`

var InitShort = `Initialize a package with the information needed for inventory tracking.`
var InitLong = `
  kpt live init [PKG_PATH] [flags]

Args:

  PKG_PATH:
    Path to the local package which should be updated with inventory information.
    It must contain a Kptfile. Defaults to the current working directory.

Flags:

  --force:
    Forces the inventory values to be updated, even if they are already set.
    Defaults to false.
  
  --inventory-id:
    Inventory identifier for the package. This is used to detect overlap between
    packages that might use the same name and namespace for the inventory object.
    Defaults to an auto-generated value.
  
  --name:
    The name for the ResourceGroup resource that contains the inventory
    for the package. Defaults to the name of the package.
  
  --namespace:
    The namespace for the ResourceGroup resource that contains the inventory
    for the package. If not provided, kpt will check if all the resources
    in the package belong in the same namespace. If they do, that namespace will
    be used. If they do not, the namespace in the user's context will be chosen.
  
  --rg-file:
    The name used for the file created for the ResourceGroup CR. Defaults to
    'resourcegroup.yaml'.
`
var InitExamples = `
  # initialize a package in the current directory.
  $ kpt live init

  # initialize a package with explicit namespace for the ResourceGroup.
  $ kpt live init --namespace=test my-dir
`

var InstallResourceGroupShort = `Install the ResourceGroup CRD in the cluster.`
var InstallResourceGroupLong = `
  kpt live install-resource-group
`
var InstallResourceGroupExamples = `
  # install ResourceGroup CRD into the current cluster.
  $ kpt live install-resource-group
`

var MigrateShort = `Migrate a package and the inventory object to use the ResourceGroup CRD.`
var MigrateLong = `
  kpt live migrate [PKG_PATH] [flags]

Args:

  PKG_PATH:
    Path to the local package. It must have a Kptfile and inventory metadata
    in the package in either the ConfigMap, Kptfile or ResourceGroup format.
    It defaults to the current directory.

Flags:

  --dry-run:
    Go through the steps of migration, but don't make any changes.
  
  --force:
    Forces the inventory values in the ResourceGroup manfiest to be updated,
    even if they are already set. Defaults to false.
  
  --name:
    The name for the ResourceGroup resource that contains the inventory
    for the package. Defaults to the same name as the existing inventory
    object.
  
  --namespace:
    The namespace for the ResourceGroup resource that contains the inventory
    for the package. If not provided, it defaults to the same namespace as the
    existing inventory object.
`
var MigrateExamples = `
  # Migrate the package in the current directory.
  $ kpt live migrate
`

var StatusShort = `Display shows the status for the resources in the cluster`
var StatusLong = `
  kpt live status [PKG_PATH | -] [flags]

Args:

  PKG_PATH | -:
    Path to the local package for which the status of the package in the cluster
    should be displayed. It must contain either a Kptfile or a ResourceGroup CR
    with inventory metadata.
    Defaults to the current working directory.
    Using '-' as the package path will cause kpt to read resources from stdin.

Flags:

  --output:
    Determines the output format for the status information. Must be one of the following:
  
      * events: The output will be a list of the status events as they become available.
      * json: The output will be a list of the status events as they become available,
        each formatted as a json object.
      * table: The output will be presented as a table that will be updated inline
        as the status of resources become available.
  
    The default value is ‘events’.
  
  --poll-period:
    The frequency with which the cluster will be polled to determine the status
    of the applied resources. The default value is 2 seconds.
  
  --poll-until:
    When to stop polling for status and exist. Must be one of the following:
  
      * known: Exit when the status for all resources have been found.
      * current: Exit when the status for all resources have reached the Current status.
      * deleted: Exit when the status for all resources have reached the NotFound
        status, i.e. all the resources have been deleted from the live state.
      * forever: Keep polling for status until interrupted.
  
    The default value is ‘known’.
  
  --timeout:
    Determines how long the command should run before exiting. This deadline will
    be enforced regardless of the value of the --poll-until flag. The default is
    to wait forever.
    
  --inv-type:
    Ways to get the inventory information. Must be one of the following:
    
    * local: Get the inventory information from the local file.
      This will only get the inventory information of the package at the given/default path.
    * remote: Get the inventory information by calling List API to the cluster.
      This will retrieve a list of inventory information from the cluster.
    
    The default value is ‘local’.
    
  --inv-names:
    Filter for printing statuses of packages with specified inventory names.
    For multiple inventory names, use comma to them.
    This must be used with --inv-type=remote.
    
  --namespaces:
    Filter for printing statuses of packages under specified namespaces.
    For multiple namespaces, use comma to separate them.
    
  --statuses:
    Filter for printing packages with specified statuses.
    For multiple statuses, use comma to separate them.
`
var StatusExamples = `
  # Monitor status for the resources belonging to the package in the current
  # directory. Wait until all resources have reconciled.
  $ kpt live status

  # Monitor status for the resources belonging to the package in the my-app
  # directory. Output in table format:
  $ kpt live status my-app --poll-until=forever --output=table

  # Monitor status for the all resources on the cluster
  # with certain inventory names and under certain namespaces.
  $ kpt live status --inv-type remote --inv-names inv1,inv2 --namespaces ns1,ns2

  # Monitor resources on the cluster that has Current or InProgress status
  $ kpt live status --inv-type remote --statuses Current,InProgress
`
