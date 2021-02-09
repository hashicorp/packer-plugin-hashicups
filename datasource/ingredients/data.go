//go:generate mapstructure-to-hcl2 -type DatasourceOutput,Config
package ingredients

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/sylviamoss/hashicups-client-go"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`

	Coffee string `mapstructure:"coffee" required:"true"`
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

	var host *string
	var username *string
	var password *string

	if d.config.Host != "" {
		host = &d.config.Host
	}
	if d.config.Username != "" {
		username = &d.config.Username
	}
	if d.config.Password != "" {
		password = &d.config.Password
	}

	client, err := hashicups.NewClient(host, username, password)
	if err != nil {
		return cty.EmptyObjectVal, err
	}

	coffees, err := client.GetCoffees()
	if err != nil {
		return cty.EmptyObjectVal, err
	}

	coffeeId := ""
	for _, coffee := range coffees {
		if coffee.Name == d.config.Coffee {
			coffeeId = strconv.Itoa(coffee.ID)
			continue
		}
	}

	if coffeeId == "" {
		return cty.EmptyObjectVal, fmt.Errorf("%s not found", d.config.Coffee)
	}

	ingredients, err := client.GetCoffeeIngredients(coffeeId)
	if err != nil {
		return cty.EmptyObjectVal, err
	}

	mapOfIngredients := map[string]string{}
	for _, ingredient := range ingredients {
		mapOfIngredients[ingredient.Name] = strconv.Itoa(ingredient.ID)
	}
	output.Map = mapOfIngredients
	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
