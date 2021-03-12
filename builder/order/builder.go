package order

import (
	"context"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

const BuilderId = "hashicups.order.builder"

type Builder struct {
	config Config
	runner multistep.Runner
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMapstructure().HCL2Spec() }

func (b *Builder) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	if err = config.Decode(&b.config, nil, raws...); err != nil {
		return nil, nil, err
	}

	// TODO check if configuration is correct

	// Let Packer know that this builder will generate a OrderId
	buildGeneratedData := []string{"OrderId"}
	return buildGeneratedData, nil, err
}

func (b *Builder) Run(ctx context.Context, ui packersdk.Ui, hook packersdk.Hook) (packersdk.Artifact, error) {
	steps := []multistep.Step{}

	// Setup the steps for this builder
	steps = append(steps,
		&StepCreateClient{Auth: b.config.AuthConfig},
		&StepCreateOrder{Items: b.config.Item},
		&StepWaitForBarista{},
		// StepProvision will run all provisioners defined in the Packer template to run with this builder
		new(commonsteps.StepProvision),
	)

	// Setup the state bag and initial state for the steps
	// The state bag is used to share information among the steps and to keep errors.
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
			"order":  state.Get("order"),
			"client": state.Get("client"),
			// Add the builder generated data to the artifact StateData so that post-processors
			// can access them.
			"generated_data": state.Get("generated_data"),
		},
	}
	return artifact, nil
}
