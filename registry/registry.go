package registry

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/sharkyze/waypoint-plugin-archive/builder"
	"google.golang.org/api/cloudfunctions/v1"
)

type RegistryConfig struct {
	// Project is the project to deploy to.
	Project string `hcl:"project,attr"`

	// Location	represents the Google Cloud location where the application
	// will be deployed, e.g. us-west1.
	Location string `hcl:"location,attr"`

	// Unauthenticated, if set to true, will allow unauthenticated access
	// to your deployment. This defaults to true.
	Unauthenticated *bool `hcl:"unauthenticated,optional"`
}

type Registry struct {
	config RegistryConfig
}

// Config implements component.Configurable.
func (r *Registry) Config() (interface{}, error) {
	return &r.config, nil
}

// Implement ConfigurableNotify
func (r *Registry) ConfigSet(config interface{}) error {
	c, ok := config.(*RegistryConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *RegisterConfig as parameter")
	}

	// validate the config
	_ = c

	return nil
}

// Implement Registry
func (r *Registry) PushFunc() interface{} {
	// return a function which will be called by Waypoint
	return r.push
}

// A PushFunc does not have a strict signature, you can define the parameters
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
//
// In addition to default input parameters the builder.Binary from the Build step
// can also be injected.
//
// The output parameters for PushFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (r *Registry) push(
	ctx context.Context,
	log hclog.Logger,
	ui terminal.UI,
	archive *builder.Archive,
) (*Artifact, error) {
	st := ui.Status()
	defer st.Close()

	st.Update("Pushing archive to Google Cloud Storage")

	artifact := Artifact{
		Project:  r.config.Project,
		Location: r.config.Location,
	}

	cloudfunctionsService, err := cloudfunctions.NewService(ctx)
	if err != nil {
		return nil, err
	}

	uploadURLCall := cloudfunctionsService.Projects.Locations.Functions.GenerateUploadUrl(
		fmt.Sprintf("projects/%s/locations/%s", r.config.Project, r.config.Location),
		&cloudfunctions.GenerateUploadUrlRequest{},
	)
	uploadURLCall = uploadURLCall.Context(ctx)

	uploadURLResponse, err := uploadURLCall.Do()
	if err != nil {
		return nil, err
	}

	uploadURL := uploadURLResponse.UploadUrl
	artifact.Source = uploadURL

	file, err := os.Open(archive.OutputPath)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	mb100 := int64(1e+8)
	if fi.Size() > mb100 {
		return nil, errors.New("File size should not exceed 100MB")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uploadURL, file)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/zip")
	req.Header.Add("x-goog-content-length-range", "0,104857600")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status + "\n\n" + string(body))
	}

	st.Step(terminal.StatusOK, "Cloud Function Archive successfully uploaded to Google Cloud Functions")

	return &artifact, nil
}
