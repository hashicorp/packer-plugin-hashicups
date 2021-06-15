//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package toppings

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

type Config struct {
	Toppings []string `mapstructure:"toppings" required:"true"`
}

type Provisioner struct {
	config Config
}

func (p *Provisioner) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, nil, raws...)
	if err != nil {
		return err
	}

	if len(p.config.Toppings) == 0 {
		return fmt.Errorf("you must specify one or more toppings")
	}
	return nil
}

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator, _ map[string]interface{}) error {
	ui.Say("Pouring some toppings")
	for _, topping := range p.config.Toppings {
		ui.Say(fmt.Sprintf("* Pouring %s...", topping))
		time.Sleep(5 * time.Second)
	}
	return nil
}
