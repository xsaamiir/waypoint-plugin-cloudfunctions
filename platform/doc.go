package platform

import (
	"github.com/hashicorp/waypoint-plugin-sdk/docs"
)

func (p *Platform) Documentation() (*docs.Documentation, error) {
	doc, err := docs.New(docs.FromConfig(&DeployConfig{}))
	if err != nil {
		return nil, err
	}

	doc.Description("Deploy a Google Cloud Function using a zip archive previously uploaded to Cloud Storage")

	doc.Example(`
project = "examples"

app "helloworld" {
  path = "./helloworld"

  build {
    use "archive" {}

    registry {
      use "cloudfunctions" {
        project = "project-id"
        location = "europe-west1"
      }
    }
  }

  deploy {
    use "cloudfunctions" {
      entry_point = "HelloHTTP"
      description = "Deployed using Waypoint 🎉"
      runtime = "go113"
      max_instances = 1
      available_memory_mb = 128
      ingress_settings = "ALLOW_ALL"
      trigger_http = true
    }
  }

  release {
    use "cloudfunctions" {
      unauthenticated = true
    }
  }
}
`)

	_ = doc.SetField(
		"environment_variables",
		"Environment Variables that shall be available during function execution.",
	)

	_ = doc.SetField(
		"build_environment_variables",
		"Build Environment Variables that shall be available during build time.",
	)

	_ = doc.SetField(
		"max_instances",
		`MaxInstances sets the maximum number of instances for the function.
A function execution that would exceed max-instances times out.`,
	)

	_ = doc.SetField(
		"runtime",
		`Runtime in which to run the function.
Available runtimes:
 - nodejs10: Node.js 10
 - nodejs12: Node.js 12
 - python37: Python 3.7
 - python38: Python 3.8
 - go111: Go 1.11
 - go113: Go 1.13
 - java11: Java 11
 - nodejs6: Node.js 6 (deprecated)
 - nodejs8: Node.js 8 (deprecated)`,
	)

	_ = doc.SetField(
		"timeout",
		`Timeout is execution timeout. 
Execution is considered failed and can be terminated if the function is not completed at the end
of the timeout period. Defaults to 60 seconds.
A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
	)

	_ = doc.SetField(
		"available_memory_mb",
		`AvailableMemoryMB is the limit on the amount of memory the function can use.
Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
By default, a new function is limited to 256MB of memory.`,
	)

	_ = doc.SetField(
		"entry_point",
		`EntryPoint is the name of the function (as defined in source code) that will be executed.
Defaults to the resource name suffix, if not specified.
For backward compatibility, if function with given name is not found, 
then the system will try to use function named "function".
For Node.js this is name of a function exported by the module specified in source_location.`,
	)

	_ = doc.SetField(
		"ingress_settings",
		`IngressSettings: The ingress settings for the function, controlling what traffic can reach it.
Possible values:
 - "INGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
 - "ALLOW_ALL" - Allow HTTP traffic from public and private sources.
 - "ALLOW_INTERNAL_ONLY" - Allow HTTP traffic from only private VPC sources.
 - "ALLOW_INTERNAL_AND_GCLB" - Allow HTTP traffic from private VPC sources and through GCLB.`,
	)

	_ = doc.SetField(
		"description",
		`Description is a user-provided description of a function.`,
	)

	_ = doc.SetField(
		"trigger_http",
		`TriggerHTTP allows any HTTP request (of a supported type) to the endpoint to trigger function execution.
Cannot be used with EventTrigger.`,
	)

	_ = doc.SetField(
		"event_trigger",
		`EventTrigger is  the source that fires events in response to a condition in another service.
Cannot be used with TriggerHTTP.`,
	)

	_ = doc.SetField("labels", `Labels associated with this Cloud Function.`)

	_ = doc.SetField(
		"network",
		`Network: The VPC Network that this cloud function can connect to.
It can be either the fully-qualified URI, or the short name of the network resource. 
If the short network name is used, the network must belong to the same project.
Otherwise,  it must belong to a project within the same organization.
The format of this field is either 'projects/{project}/global/networks/{network}' 
or '{network}', where {project} is a project id where the network is defined,
and {network} is the short name of the network.
This field is mutually exclusive with 'vpc_connector' and will be replaced by it. 
See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more
information on connecting Cloud projects.`,
	)

	_ = doc.SetField(
		"vpc_connector",
		`VpcConnector: The VPC Network Connector that this cloud function can connect to. 
It can be either the fully-qualified URI, or the short name of the network connector resource. 
The format of this field is 'projects/*/locations/*/connectors/*'.
This field is mutually exclusive with 'network' field and will eventually replace it.
See	[the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information
on connecting Cloud projects.`,
	)

	_ = doc.SetField(
		"vpc_connector_egress_settings",
		`The egress settings for the connector, controlling what traffic is diverted through it.
Possible values:
 - "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
 - "PRIVATE_RANGES_ONLY" - Use the VPC Access Connector only for private IP space from RFC1918.
 - "ALL_TRAFFIC" - Force the use of VPC Access Connector for all egress traffic from the function.`,
	)

	return doc, nil
}
