package release

import (
	"context"
	"fmt"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"google.golang.org/api/cloudfunctions/v1"

	"github.com/sharkyze/waypoint-plugin-cloudfunctions/platform"
)

type ReleaseConfig struct {
	// Unauthenticated, if set to true, will allow unauthenticated access
	// to your deployment. This defaults to false.
	Unauthenticated bool `hcl:"unauthenticated,optional"`
}

type ReleaseManager struct {
	config ReleaseConfig
}

// Config implements component.Configurable.
func (rm *ReleaseManager) Config() (interface{}, error) {
	return &rm.config, nil
}

// ConfigSet implements component.ConfigurableNotify.
func (rm *ReleaseManager) ConfigSet(config interface{}) error {
	_, ok := config.(*ReleaseConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *ReleaseConfig as parameter")
	}

	// validate the config

	return nil
}

// ReleaseFunc implements component.Builder.
func (rm *ReleaseManager) ReleaseFunc() interface{} {
	// return a function which will be called by Waypoint
	return rm.release
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

// In addition to default input parameters the platform.Deployment from the Deploy step
// can also be injected.
//
// The output parameters for ReleaseFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
//
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (rm *ReleaseManager) release(
	ctx context.Context,
	ui terminal.UI,
	deployment *platform.Deployment,
) (*Release, error) {
	st := ui.Status()
	defer st.Close()

	release := Release{
		Version: deployment.Version,
		Name:    deployment.Name,
		Url:     deployment.Url,
	}

	// TODO: handle cases where the parameter has changed
	// 	from Authenticated to Unauthenticated ?
	if !rm.config.Unauthenticated {
		st.Step(
			terminal.StatusOK,
			"No Operation release, Cloud Function already deployed but only accessible to authenticated users",
		)
		return &release, nil
	}

	st.Update("Releasing Google Cloud Function to all unauthenticated users")

	cloudfunctionsService, err := cloudfunctions.NewService(ctx)
	if err != nil {
		st.Step(terminal.StatusError, "Error setting IAM Policy to allUsers")
		return nil, err
	}

	err = setIAMPolicyAllUsers(ctx, cloudfunctionsService, release.Name)
	if err != nil {
		return nil, err
	}

	st.Step(terminal.StatusOK, "IAM Policy successfully set to 'allUsers'")

	return &release, nil
}

// setIAMPolicyAllUsers sets the IAM policy on the deployment so that anyone
// can access it (no auth required).
func setIAMPolicyAllUsers(
	ctx context.Context,
	service *cloudfunctions.Service,
	name string,
) error {
	policyRequest := cloudfunctions.SetIamPolicyRequest{
		Policy: &cloudfunctions.Policy{
			Bindings: []*cloudfunctions.Binding{
				{
					Role:    "roles/cloudfunctions.invoker",
					Members: []string{"allUsers"},
				},
			},
		},
	}

	_, err := service.Projects.Locations.Functions.SetIamPolicy(name, &policyRequest).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

// URL implements component.Release.
func (x *Release) URL() string { return x.Url }
