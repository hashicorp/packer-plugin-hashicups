// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"packer-plugin-hashicups/builder/order"
	"packer-plugin-hashicups/datasource/coffees"
	"packer-plugin-hashicups/datasource/ingredients"
	"packer-plugin-hashicups/post-processor/receipt"
	"packer-plugin-hashicups/provisioner/toppings"
	"packer-plugin-hashicups/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("coffees", new(coffees.Datasource))
	pps.RegisterDatasource("ingredients", new(ingredients.Datasource))
	pps.RegisterBuilder("order", new(order.Builder))
	pps.RegisterProvisioner("toppings", new(toppings.Provisioner))
	pps.RegisterPostProcessor("receipt", new(receipt.PostProcessor))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
