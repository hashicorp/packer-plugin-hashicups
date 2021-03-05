package order

import (
	"context"
	"packer-plugin-hashicups/common"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepCreateClient struct {
	Auth common.AuthConfig
}

func (s *StepCreateClient) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	ui.Say("Creating HashiCups Client")
	client, err := s.Auth.CreateClient()
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	state.Put("client", client)
	return multistep.ActionContinue
}

func (s *StepCreateClient) Cleanup(state multistep.StateBag) {
	// Nothing to clean
}
