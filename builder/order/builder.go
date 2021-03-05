//go:generate mapstructure-to-hcl2 -type Config,OrderItem,Coffee,Ingredient

package order

import (
	"context"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

const BuilderId = "hashicups.builder"

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	Username            string      `mapstructure:"username"`
	Password            string      `mapstructure:"password"`
	Host                string      `mapstructure:"host"`
	Item                []OrderItem `mapstructure:"item,omitempty" required:"true"`
}

type OrderItem struct {
	Coffee   Coffee `mapstructure:"coffee" required:"true"`
	Quantity int    `mapstructure:"quantity" required:"true"`
}

type Coffee struct {
	ID         string       `mapstructure:"id" required:"true"`
	Name       string       `mapstructure:"name" required:"true"`
	Ingredient []Ingredient `mapstructure:"ingredient"`
}

type Ingredient struct {
	ID       string `mapstructure:"id" required:"true"`
	Quantity int    `mapstructure:"quantity"`
}

type Builder struct {
	config Config
	runner multistep.Runner
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMapstructure().HCL2Spec() }

func (b *Builder) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	buildGeneratedData := []string{"OrderId"}
	return buildGeneratedData, nil, config.Decode(&b.config, nil, raws...)
}

func (b *Builder) Run(ctx context.Context, ui packersdk.Ui, hook packersdk.Hook) (packersdk.Artifact, error) {
	steps := []multistep.Step{}

	steps = append(steps,
		&StepCreateClient{
			Username: b.config.Username,
			Password: b.config.Password,
			Host:     b.config.Host,
		},
		&StepCreateOrder{Items: b.config.Item},
		new(commonsteps.StepProvision),
	)

	// Setup the state bag and initial state for the steps
	state := new(multistep.BasicStateBag)
	state.Put("hook", hook)
	state.Put("ui", ui)

	// Run!
	b.runner = commonsteps.NewRunner(steps, b.config.PackerConfig, ui)
	b.runner.Run(ctx, state)

	// If there was an error, return that
	if err, ok := state.GetOk("error"); ok {
		return nil, err.(error)
	}

	artifact := &Artifact{
		StateData: map[string]interface{}{
			"order": state.Get("order"),
			"client": state.Get("client"),
			// Add the builder generated data to the artifact StateData so that post-processors
			// can access them.
			"generated_data": state.Get("generated_data"),
		},
	}
	return artifact, nil
}
