//go:generate mapstructure-to-hcl2 -type DatasourceOutput,Config
package coffees

import (
	"packer-plugin-hashicups/common"
	"strconv"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	common.AuthConfig `mapstructure:",squash"`
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
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	output := DatasourceOutput{}

	client, err := d.config.AuthConfig.CreateClient()
	if err != nil {
		return cty.EmptyObjectVal, err
	}

	coffees, err := client.GetCoffees()
	if err != nil {
		return cty.EmptyObjectVal, err
	}

	mapOfCoffees := map[string]string{}
	for _, coffee := range coffees {
		mapOfCoffees[coffee.Name] = strconv.Itoa(coffee.ID)
	}
	output.Map = mapOfCoffees
	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
