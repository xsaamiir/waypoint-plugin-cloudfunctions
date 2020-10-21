package main

import (
	sdk "github.com/hashicorp/waypoint-plugin-sdk"

	"github.com/sharkyze/waypoint-plugin-cloudfunctions/platform"
	"github.com/sharkyze/waypoint-plugin-cloudfunctions/registry"
	"github.com/sharkyze/waypoint-plugin-cloudfunctions/release"
)

func main() {
	// sdk.Main allows you to register the components which should
	// be included in your plugin
	// Main sets up all the go-plugin requirements

	sdk.Main(sdk.WithComponents(
		// Comment out any components which are not
		// required for your plugin
		&registry.Registry{},
		&platform.Platform{},
		&release.ReleaseManager{},
	))
}
