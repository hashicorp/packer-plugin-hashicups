//go:generate mapstructure-to-hcl2 -type Config

package status

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/schollz/progressbar/v3"
)

type Config struct {
	OrderId string `mapstructure:"order"`
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
	return nil
}

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator, _ map[string]interface{}) error {
	ui.Say(fmt.Sprintf("Waiting for order %q status", p.config.OrderId))
	bar := progressbar.Default(10000)

	for i := 0; i < 10000; i++ {
		bar.Add(1)
		time.Sleep(time.Millisecond)
	}
	return nil
}
