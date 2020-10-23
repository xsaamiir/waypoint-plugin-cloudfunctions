package release

import (
	"github.com/hashicorp/waypoint-plugin-sdk/docs"
)

func (rm *ReleaseManager) Documentation() (*docs.Documentation, error) {
	doc, err := docs.New(docs.FromConfig(&ReleaseConfig{}))
	if err != nil {
		return nil, err
	}

	doc.Description("If the function is unauthenticated, set the required IAM Policies. Otherwise a No-operation.")

	_ = doc.SetField(
		"unauthenticated",
		`If set to true, will allow unauthenticated access to your deployment. This defaults to false.`,
	)

	return doc, nil
}
