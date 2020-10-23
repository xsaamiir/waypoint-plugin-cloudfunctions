package platform

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/googleapi"

	"github.com/sharkyze/waypoint-plugin-cloudfunctions/internal/cloudfunctionsutil"
	"github.com/sharkyze/waypoint-plugin-cloudfunctions/registry"
)

type DeployConfig struct {
	// EnvironmentVariables that shall be available during function execution.
	EnvironmentVariables map[string]string `hcl:"environment_variables,optional"`

	// BuildEnvironmentVariables that shall be available during build time.
	BuildEnvironmentVariables map[string]string `hcl:"build_environment_variables,optional"`

	// MaxInstances sets the maximum number of instances for the function.
	// A function execution that would exceed max-instances times out.
	MaxInstances int64 `hcl:"max_instances,optional"`

	// Runtime in which to run the function.
	// Available runtimes:
	// 	nodejs10: Node.js 10
	// 	nodejs12: Node.js 12
	// 	python37: Python 3.7
	// 	python38: Python 3.8
	// 	go111: Go 1.11
	// 	go113: Go 1.13
	// 	java11: Java 11
	// 	nodejs6: Node.js 6 (deprecated)
	// 	nodejs8: Node.js 8 (deprecated)
	Runtime string `hcl:"runtime,optional"`

	// Timeout is execution timeout. Execution is considered failed and can be terminated if the function is not
	// completed at the end of the timeout period. Defaults to 60 seconds.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	Timeout string `hcl:"timeout,optional"`

	// AvailableMemoryMB is the limit on the amount of memory the function can use.
	// Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.
	// By default, a new function is limited to 256MB of memory.
	AvailableMemoryMB int64 `hcl:"available_memory_mb,optional"`

	// EntryPoint is the name of the function (as defined in source code) that will be executed.
	// Defaults to the resource name suffix, if not specified.
	// For backward compatibility, if function with given name is not found, then the system will try to use function named "function".
	// For Node.js this is name of a function exported by the module specified in source_location.
	EntryPoint string `hcl:"entry_point,optional"`

	// IngressSettings: The ingress settings for the function, controlling
	// what traffic can reach it.
	//
	// Possible values:
	//   "INGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
	//   "ALLOW_ALL" - Allow HTTP traffic from public and private sources.
	//   "ALLOW_INTERNAL_ONLY" - Allow HTTP traffic from only private VPC
	// sources.
	//   "ALLOW_INTERNAL_AND_GCLB" - Allow HTTP traffic from private VPC
	// sources and through GCLB.
	IngressSettings string `hcl:"ingress_settings,optional"`

	// Description is a user-provided description of a function.
	Description string `hcl:"description,optional"`

	// TriggerHTTP allows any HTTP request (of a supported type) to the
	// endpoint to trigger function execution.
	// Cannot be used with EventTrigger.
	TriggerHTTP triggerHTTP `hcl:"trigger_http,optional"`

	// EventTrigger is  the source that fires events in response to a condition
	// in another service.
	// Cannot be used with TriggerHTTP.
	EventTrigger *eventTrigger `hcl:"event_trigger,block"`

	// Labels: Labels associated with this Cloud Function.
	Labels map[string]string `hcl:"labels,optional"`

	// Network: The VPC Network that this cloud function can connect to. It
	// can be either the fully-qualified URI, or the short name of the
	// network resource. If the short network name is used, the network must
	// belong to the same project. Otherwise, it must belong to a project
	// within the same organization. The format of this field is either
	// `projects/{project}/global/networks/{network}` or `{network}`, where
	// {project} is a project id where the network is defined, and {network}
	// is the short name of the network. This field is mutually exclusive
	// with `vpc_connector` and will be replaced by it. See [the VPC
	// documentation](https://cloud.google.com/compute/docs/vpc) for more
	// information on connecting Cloud projects.
	Network string `hcl:"network,optional"`

	// VpcConnector: The VPC Network Connector that this cloud function can
	// connect to. It can be either the fully-qualified URI, or the short
	// name of the network connector resource. The format of this field is
	// `projects/*/locations/*/connectors/*` This field is mutually
	// exclusive with `network` field and will eventually replace it. See
	// [the VPC documentation](https://cloud.google.com/compute/docs/vpc)
	// for more information on connecting Cloud projects.
	VpcConnector string `hcl:"vpc_connector,optional"`

	// VpcConnectorEgressSettings: The egress settings for the connector,
	// controlling what traffic is diverted through it.
	//
	// Possible values:
	//   "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED" - Unspecified.
	//   "PRIVATE_RANGES_ONLY" - Use the VPC Access Connector only for
	// private IP space from RFC1918.
	//   "ALL_TRAFFIC" - Force the use of VPC Access Connector for all
	// egress traffic from the function.
	VpcConnectorEgressSettings string `hcl:"vpc_connector_egress_settings,optional"`
}

type eventTrigger struct {
	// EventType: Required. The type of event to observe. For example:
	// `providers/cloud.storage/eventTypes/object.change` and
	// `providers/cloud.pubsub/eventTypes/topic.publish`. Event types match
	// pattern `providers/*/eventTypes/*.*`. The pattern contains:
	// 	1. namespace: For example, `cloud.storage` and `google.firebase.analytics`.
	// 	2. resource type: The type of resource on which event occurs.
	// 	For example, the Google Cloud Storage API includes the type `object`.
	// 	3. action: The action that generates the event.
	// 	For example, action for a Google Cloud Storage Object is 'change'.
	// 	These parts are lower case.
	EventType string `hcl:"event_type"`

	// Resource: Required. The resource(s) from which to observe events, for
	// example, `projects/project-name/buckets/myBucket`.
	// Not all syntactically correct values are accepted by all services.
	// For example:
	// 	1. The authorization model must support it. Google Cloud Functions only
	//  allows EventTriggers to be deployed that observe resources in the
	//  same project as the `CloudFunction`.
	// 	2. The resource type must match  the pattern expected for an `event_type`.
	// 	For example, an `EventTrigger` that has an `event_type` of
	// 	"google.pubsub.topic.publish" should have a resource that matches
	// 	Google Cloud Pub/Sub topics.
	// Additionally, some services may support short names when creating an `EventTrigger`.
	// These will always be  returned in the normalized "long" format. See each *service's*
	// documentation for supported formats.
	Resource string `hcl:"resource"`

	// FailurePolicy: Specifies policy for failed executions.
	FailurePolicy failurePolicy `hcl:"failure_policy,block"`

	// Service: The hostname of the service that should be observed. If no
	// string is provided, the default service implementing the API will be
	// used. For example, `storage.googleapis.com` is the default for all
	// event types in the `google.storage` namespace.
	Service string `hcl:"service,optional"`
}

func (t *eventTrigger) toCF() *cloudfunctions.EventTrigger {
	if t == nil {
		return nil
	}

	return &cloudfunctions.EventTrigger{
		EventType:     t.EventType,
		FailurePolicy: t.FailurePolicy.toCF(),
		Resource:      t.Resource,
		Service:       t.Service,
	}
}

type failurePolicy struct {
	// Retry: Describes the retry policy in case of function's execution
	// failure. A function execution will be retried on any failure. A
	// failed execution will be retried up to 7 days with an exponential
	// backoff (capped at 10 seconds). Retried execution is charged as any
	// other execution.
	Retry bool `hcl:"retry"`
}

func (p *failurePolicy) toCF() *cloudfunctions.FailurePolicy {
	if p == nil || !p.Retry {
		return nil
	}

	return &cloudfunctions.FailurePolicy{Retry: &cloudfunctions.Retry{}}
}

type triggerHTTP bool

// toCF returns the data in the format expected by the cloudfunctions.Service.
func (t triggerHTTP) toCF() *cloudfunctions.HttpsTrigger {
	if !t {
		return nil
	}

	return &cloudfunctions.HttpsTrigger{}
}

type Platform struct {
	config DeployConfig
}

// Config implements component.Configurable.
func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

// ConfigSet implements component.ConfigurableNotify.
func (p *Platform) ConfigSet(config interface{}) error {
	c, ok := config.(*DeployConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *DeployConfig as parameter")
	}

	// validate the config
	if c.TriggerHTTP && c.EventTrigger != nil {
		return fmt.Errorf("trigger_http and event_type cannot be used together")
	}

	return nil
}

// DeployFunc implements component.Builder.
func (p *Platform) DeployFunc() interface{} {
	// return a function which will be called by Waypoint
	return p.deploy
}

// A BuildFunc does not have a strict signature, you can define the parameters
// you need based on the Available parameters that the Waypoint SDK provides.
// Waypoint will automatically inject parameters as specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - *datadir.Project
// - *datadir.App
// - *datadir.Component
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet

// In addition to default input parameters the registry.Artifact from the Build step
// can also be injected.
//
// The output parameters for BuildFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (p *Platform) deploy(
	ctx context.Context,
	source *component.Source,
	ui terminal.UI,
	artifact *registry.Artifact,
) (*Deployment, error) {
	st := ui.Status()
	defer st.Close()

	project := artifact.Project
	location := artifact.Location
	sourceApp := source.App
	functionName := fmt.Sprintf("projects/%s/locations/%s/functions/%s", project, location, sourceApp)

	st.Update("Deploying Google Cloud Function '" + functionName + "'")

	cloudfunctionsService, err := cloudfunctions.NewService(ctx)
	if err != nil {
		return nil, err
	}

	// We need to determine if we're creating or updating a function. To
	// do this, we just query GCP directly.
	create := false
	getCall := cloudfunctionsService.Projects.Locations.Functions.Get(functionName).Context(ctx)

	st.Update("Checking if function already exists " + functionName + "'")

	var gerr *googleapi.Error
	cf, err := getCall.Do()
	if err != nil {
		if errors.As(err, &gerr) && gerr.Code == 404 {
			create = true
		} else {
			st.Step(terminal.StatusError, "Error fetching function")
			return nil, err
		}
	}

	if create {
		st.Step(terminal.StatusOK, "Google Cloud Function does not exist, creating function")
	} else {
		st.Step(terminal.StatusOK, "Google Cloud Function already exists, updating function")
	}

	var op *cloudfunctions.Operation

	if create {
		cloudFuncReq := cloudfunctions.CloudFunction{
			AvailableMemoryMb:          p.config.AvailableMemoryMB,
			BuildEnvironmentVariables:  p.config.BuildEnvironmentVariables,
			Description:                p.config.Description,
			EntryPoint:                 p.config.EntryPoint,
			EnvironmentVariables:       p.config.EnvironmentVariables,
			EventTrigger:               p.config.EventTrigger.toCF(),
			HttpsTrigger:               p.config.TriggerHTTP.toCF(),
			IngressSettings:            p.config.IngressSettings,
			Labels:                     p.config.Labels,
			MaxInstances:               p.config.MaxInstances,
			Name:                       functionName,
			Network:                    p.config.Network,
			Runtime:                    p.config.Runtime,
			ServiceAccountEmail:        "",
			SourceArchiveUrl:           "",
			SourceRepository:           nil,
			SourceUploadUrl:            artifact.Source,
			Timeout:                    p.config.Timeout,
			VpcConnector:               p.config.VpcConnector,
			VpcConnectorEgressSettings: p.config.VpcConnectorEgressSettings,
		}

		op, err = createFunction(ctx, cloudfunctionsService, project, location, &cloudFuncReq)
		if err != nil {
			st.Step(terminal.StatusError, fmt.Sprintf("%#v", *err.(*googleapi.Error)))
			time.Sleep(15 * time.Second)
			return nil, err
		}
	} else {
		// TODO: handle any other updated fields passed as parameters to waypoint.
		cf.SourceUploadUrl = artifact.Source
		op, err = patchFunc(ctx, cloudfunctionsService, cf)
		if err != nil {
			st.Step(terminal.StatusError, fmt.Sprintf("Error updating function: %#v", *err.(*googleapi.Error)))
			return nil, err
		}
	}

	st.Update("Building Function '" + op.Name + "'")

	op, err = cloudfunctionsutil.WaitForOperation(ctx, cloudfunctionsService, op)
	if err != nil {
		st.Step(terminal.StatusError, "Error fetching build status")
		return nil, err
	}

	if op.Error != nil {
		st.Step(terminal.StatusError, "Build error")
		return nil, errors.New(op.Error.Message)
	}

	var cfresp cloudfunctions.CloudFunction
	err = json.Unmarshal(op.Response, &cfresp)
	if err != nil {
		st.Step(terminal.StatusError, "Error reading the response data but function successfully deployed")
		return nil, err
	}

	versionID := cfresp.VersionId

	st.Step(terminal.StatusOK, fmt.Sprintf("Google Cloud Function successfully deployed 'v%d'", versionID))

	var url string
	if t := cfresp.HttpsTrigger; t != nil {
		url = t.Url
	}

	return &Deployment{Name: cfresp.Name, Version: versionID, Url: url}, nil
}

func createFunction(
	ctx context.Context,
	service *cloudfunctions.Service,
	project, location string,
	req *cloudfunctions.CloudFunction,
) (*cloudfunctions.Operation, error) {
	call := service.Projects.Locations.Functions.
		Create(
			fmt.Sprintf("projects/%s/locations/%s", project, location),
			req,
		).
		Context(ctx)

	op, err := call.Do()
	if err != nil {
		return nil, err
	}

	return op, nil
}

func patchFunc(
	ctx context.Context,
	service *cloudfunctions.Service,
	req *cloudfunctions.CloudFunction,
) (*cloudfunctions.Operation, error) {
	patchCall := service.Projects.Locations.Functions.
		Patch(req.Name, req).
		Context(ctx).
		UpdateMask("sourceUploadUrl")

	op, err := patchCall.Do()
	if err != nil {
		return nil, err
	}

	return op, nil
}
