// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type DatasourceOutput,Config
package ingredients

import (
	"fmt"
	"packer-plugin-hashicups/common"
	"strconv"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	common.AuthConfig `mapstructure:",squash"`
	Coffee            string `mapstructure:"coffee" required:"true"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	Map map[string]string `mapstructure:"map"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	if d.config.Coffee == "" {
		return fmt.Errorf("you must specify the name of the coffee to get its ingredients")
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	output := DatasourceOutput{}
	emptyOutput := hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec())

	client, err := d.config.AuthConfig.CreateClient()
	if err != nil {
		return emptyOutput, err
	}

	coffees, err := client.GetCoffees()
	if err != nil {
		return emptyOutput, err
	}

	coffeeId := ""
	for _, coffee := range coffees {
		if coffee.Name == d.config.Coffee {
			coffeeId = strconv.Itoa(coffee.ID)
			continue
		}
	}

	if coffeeId == "" {
		return emptyOutput, fmt.Errorf("%s not found", d.config.Coffee)
	}

	ingredients, err := client.GetCoffeeIngredients(coffeeId)
	if err != nil {
		return emptyOutput, err
	}

	mapOfIngredients := map[string]string{}
	for _, ingredient := range ingredients {
		mapOfIngredients[ingredient.Name] = strconv.Itoa(ingredient.ID)
	}
	output.Map = mapOfIngredients
	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
