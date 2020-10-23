[comment]: <> (!!! AUTO GENERATED, DO NOT EDIT !!!)

## cloudfunctions (registry)

Upload the source code as a zip archive to Google Cloud Storage

### Variables

#### location
Location represents the Google Cloud location where the application will be deployed, e.g. us-west1.


* Type: **string**

#### project
Project is the project to deploy to.


* Type: **string**
## cloudfunctions (platform)

Deploy a Google Cloud Function using a zip archive previously uploaded to Cloud Storage

### Variables

#### available_memory_mb
AvailableMemoryMB is the limit on the amount of memory the function can use.
Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
By default, a new function is limited to 256MB of memory.


* Type: **int64**
* __Optional__

#### build_environment_variables
Build Environment Variables that shall be available during build time.


* Type: **map[string]string**
* __Optional__

#### description
Description is a user-provided description of a function.


* Type: **string**
* __Optional__

#### entry_point
EntryPoint is the name of the function (as defined in source code) that will be executed.
Defaults to the resource name suffix, if not specified.
For backward compatibility, if function with given name is not found, 
then the system will try to use function named "function".
For Node.js this is name of a function exported by the module specified in source_location.


* Type: **string**
* __Optional__

#### environment_variables
Environment Variables that shall be available during function execution.


* Type: **map[string]string**
* __Optional__

#### event_trigger
EventTrigger is  the source that fires events in response to a condition in another service.
Cannot be used with TriggerHTTP.


* Type: ***platform.eventTrigger**

#### ingress_settings
IngressSettings: The ingress settings for the function, controlling what traffic can reach it.
Possible values:
 - "INGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
 - "ALLOW_ALL" - Allow HTTP traffic from public and private sources.
 - "ALLOW_INTERNAL_ONLY" - Allow HTTP traffic from only private VPC sources.
 - "ALLOW_INTERNAL_AND_GCLB" - Allow HTTP traffic from private VPC sources and through GCLB.


* Type: **string**
* __Optional__

#### labels
Labels associated with this Cloud Function.


* Type: **map[string]string**
* __Optional__

#### max_instances
MaxInstances sets the maximum number of instances for the function.
A function execution that would exceed max-instances times out.


* Type: **int64**
* __Optional__

#### network
Network: The VPC Network that this cloud function can connect to.
It can be either the fully-qualified URI, or the short name of the network resource. 
If the short network name is used, the network must belong to the same project.
Otherwise,  it must belong to a project within the same organization.
The format of this field is either 'projects/{project}/global/networks/{network}' 
or '{network}', where {project} is a project id where the network is defined,
and {network} is the short name of the network.
This field is mutually exclusive with 'vpc_connector' and will be replaced by it. 
See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more
information on connecting Cloud projects.


* Type: **string**
* __Optional__

#### runtime
Runtime in which to run the function.
Available runtimes:
 - nodejs10: Node.js 10
 - nodejs12: Node.js 12
 - python37: Python 3.7
 - python38: Python 3.8
 - go111: Go 1.11
 - go113: Go 1.13
 - java11: Java 11
 - nodejs6: Node.js 6 (deprecated)
 - nodejs8: Node.js 8 (deprecated)


* Type: **string**
* __Optional__

#### timeout
Timeout is execution timeout. 
Execution is considered failed and can be terminated if the function is not completed at the end
of the timeout period. Defaults to 60 seconds.
A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".


* Type: **string**
* __Optional__

#### trigger_http
TriggerHTTP allows any HTTP request (of a supported type) to the endpoint to trigger function execution.
Cannot be used with EventTrigger.


* Type: **platform.triggerHTTP**
* __Optional__

#### vpc_connector
VpcConnector: The VPC Network Connector that this cloud function can connect to. 
It can be either the fully-qualified URI, or the short name of the network connector resource. 
The format of this field is 'projects/*/locations/*/connectors/*'.
This field is mutually exclusive with 'network' field and will eventually replace it.
See	[the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information
on connecting Cloud projects.


* Type: **string**
* __Optional__

#### vpc_connector_egress_settings
The egress settings for the connector, controlling what traffic is diverted through it.
Possible values:
 - "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
 - "PRIVATE_RANGES_ONLY" - Use the VPC Access Connector only for private IP space from RFC1918.
 - "ALL_TRAFFIC" - Force the use of VPC Access Connector for all egress traffic from the function.


* Type: **string**
* __Optional__
## cloudfunctions (releasemanager)

If the function is unauthenticated, set the required IAM Policies. Otherwise a No-operation.

### Variables

#### unauthenticated
If set to true, will allow unauthenticated access to your deployment. This defaults to false.


* Type: **bool**
* __Optional__
## cloudfunctions (registry)

Upload the source code as a zip archive to Google Cloud Storage

### Variables

#### location
Location represents the Google Cloud location where the application will be deployed, e.g. us-west1.


* Type: **string**

#### project
Project is the project to deploy to.


* Type: **string**
## cloudfunctions (platform)

Deploy a Google Cloud Function using a zip archive previously uploaded to Cloud Storage

### Variables

#### available_memory_mb
AvailableMemoryMB is the limit on the amount of memory the function can use.
Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
By default, a new function is limited to 256MB of memory.


* Type: **int64**
* __Optional__

#### build_environment_variables
Build Environment Variables that shall be available during build time.


* Type: **map[string]string**
* __Optional__

#### description
Description is a user-provided description of a function.


* Type: **string**
* __Optional__

#### entry_point
EntryPoint is the name of the function (as defined in source code) that will be executed.
Defaults to the resource name suffix, if not specified.
For backward compatibility, if function with given name is not found, 
then the system will try to use function named "function".
For Node.js this is name of a function exported by the module specified in source_location.


* Type: **string**
* __Optional__

#### environment_variables
Environment Variables that shall be available during function execution.


* Type: **map[string]string**
* __Optional__

#### event_trigger
EventTrigger is  the source that fires events in response to a condition in another service.
Cannot be used with TriggerHTTP.


* Type: ***platform.eventTrigger**

#### ingress_settings
IngressSettings: The ingress settings for the function, controlling what traffic can reach it.
Possible values:
 - "INGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
 - "ALLOW_ALL" - Allow HTTP traffic from public and private sources.
 - "ALLOW_INTERNAL_ONLY" - Allow HTTP traffic from only private VPC sources.
 - "ALLOW_INTERNAL_AND_GCLB" - Allow HTTP traffic from private VPC sources and through GCLB.


* Type: **string**
* __Optional__

#### labels
Labels associated with this Cloud Function.


* Type: **map[string]string**
* __Optional__

#### max_instances
MaxInstances sets the maximum number of instances for the function.
A function execution that would exceed max-instances times out.


* Type: **int64**
* __Optional__

#### network
Network: The VPC Network that this cloud function can connect to.
It can be either the fully-qualified URI, or the short name of the network resource. 
If the short network name is used, the network must belong to the same project.
Otherwise,  it must belong to a project within the same organization.
The format of this field is either 'projects/{project}/global/networks/{network}' 
or '{network}', where {project} is a project id where the network is defined,
and {network} is the short name of the network.
This field is mutually exclusive with 'vpc_connector' and will be replaced by it. 
See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more
information on connecting Cloud projects.


* Type: **string**
* __Optional__

#### runtime
Runtime in which to run the function.
Available runtimes:
 - nodejs10: Node.js 10
 - nodejs12: Node.js 12
 - python37: Python 3.7
 - python38: Python 3.8
 - go111: Go 1.11
 - go113: Go 1.13
 - java11: Java 11
 - nodejs6: Node.js 6 (deprecated)
 - nodejs8: Node.js 8 (deprecated)


* Type: **string**
* __Optional__

#### timeout
Timeout is execution timeout. 
Execution is considered failed and can be terminated if the function is not completed at the end
of the timeout period. Defaults to 60 seconds.
A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".


* Type: **string**
* __Optional__

#### trigger_http
TriggerHTTP allows any HTTP request (of a supported type) to the endpoint to trigger function execution.
Cannot be used with EventTrigger.


* Type: **platform.triggerHTTP**
* __Optional__

#### vpc_connector
VpcConnector: The VPC Network Connector that this cloud function can connect to. 
It can be either the fully-qualified URI, or the short name of the network connector resource. 
The format of this field is 'projects/*/locations/*/connectors/*'.
This field is mutually exclusive with 'network' field and will eventually replace it.
See	[the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information
on connecting Cloud projects.


* Type: **string**
* __Optional__

#### vpc_connector_egress_settings
The egress settings for the connector, controlling what traffic is diverted through it.
Possible values:
 - "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
 - "PRIVATE_RANGES_ONLY" - Use the VPC Access Connector only for private IP space from RFC1918.
 - "ALL_TRAFFIC" - Force the use of VPC Access Connector for all egress traffic from the function.


* Type: **string**
* __Optional__
## cloudfunctions (releasemanager)

If the function is unauthenticated, set the required IAM Policies. Otherwise a No-operation.

### Variables

#### unauthenticated
If set to true, will allow unauthenticated access to your deployment. This defaults to false.


* Type: **bool**
* __Optional__
