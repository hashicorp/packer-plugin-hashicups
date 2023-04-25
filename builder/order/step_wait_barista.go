// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package order

import (
	"context"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepWaitForBarista struct {
}

func (s *StepWaitForBarista) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	ui.Say("The barista is preparing your order...")
	ui.Say("Take a seat and relax. :)")
	time.Sleep(15 * time.Second)
	ui.Say("Your order is ready!!")

	return multistep.ActionContinue
}

func (s *StepWaitForBarista) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
