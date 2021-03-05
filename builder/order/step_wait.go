package order

import (
	"context"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"time"
)

type StepWait struct {
}

func (s *StepWait) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	ui.Say("The barista is preparing your order...")
	ui.Say("Take a seat and relax. :)")
	time.Sleep(15 * time.Second)
	ui.Say("Your order is ready!!")

	return multistep.ActionContinue
}

func (s *StepWait) Cleanup(state multistep.StateBag) {
	// Nothing to clean
}
