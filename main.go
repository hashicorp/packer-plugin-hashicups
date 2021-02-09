package main

import (
	"fmt"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
	"os"
	"packer-plugin-hashicups/builder/order"
	"packer-plugin-hashicups/datasource/coffees"
	"packer-plugin-hashicups/datasource/ingredients"
)

var (
	// Version is the main version number that is being run at the moment.
	Version = "0.0.1"

	// VersionPrerelease is A pre-release marker for the Version. If this is ""
	// (empty string) then it means that it is a final release. Otherwise, this
	// is a pre-release such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "dev"

	// PluginVersion is used by the plugin set to allow Packer to recognize
	// what version this plugin is.
	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("coffees", new(coffees.Datasource))
	pps.RegisterDatasource("ingredients", new(ingredients.Datasource))
	pps.RegisterBuilder("order", new(order.Builder))
	//pps.RegisterProvisioner("my-provisioner", new(scaffoldingProv.Provisioner))
	//pps.RegisterPostProcessor("my-post-processor", new(scaffoldingPP.PostProcessor))
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
