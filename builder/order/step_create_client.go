package order

import (
	"context"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/sylviamoss/hashicups-client-go"
)

type StepCreateClient struct {
	Username string
	Password string
	Host     string
}

func (s *StepCreateClient) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	var host *string
	var username *string
	var password *string

	if s.Host != "" {
		host = &s.Host
	}
	if s.Username != "" {
		username = &s.Username
	}
	if s.Password != "" {
		password = &s.Password
	}

	ui.Say("Creating HashiCups Client")
	client, err := hashicups.NewClient(host, username, password)
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
