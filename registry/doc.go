package registry

import (
	"github.com/hashicorp/waypoint-plugin-sdk/docs"
)

func (r *Registry) Documentation() (*docs.Documentation, error) {
	doc, err := docs.New(docs.FromConfig(&RegistryConfig{}))
	if err != nil {
		return nil, err
	}

	doc.Description("Upload the source code as a zip archive to Google Cloud Storage")

	_ = doc.SetField("project", `Project is the project to deploy to.`)

	_ = doc.SetField(
		"location",
		`Location represents the Google Cloud location where the application will be deployed, e.g. us-west1.`,
	)

	return doc, nil
}
