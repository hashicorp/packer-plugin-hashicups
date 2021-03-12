package order

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/sylviamoss/hashicups-client-go"
)

type StepCreateOrder struct {
	Items []OrderItem
}

func (s *StepCreateOrder) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	client := state.Get("client").(*hashicups.Client)

	ui.Say("Creating HashiCups Order")

	orderItems := []hashicups.OrderItem{}
	for _, item := range s.Items {

		ingredients := []hashicups.Ingredient{}
		for _, ig := range item.Coffee.Ingredient {
			id, _ := strconv.Atoi(ig.ID)
			ingredients = append(ingredients, hashicups.Ingredient{
				ID:       id,
				Quantity: ig.Quantity,
			})
		}

		id, _ := strconv.Atoi(item.Coffee.ID)
		oi := hashicups.OrderItem{
			Coffee: hashicups.Coffee{
				ID:         id,
				Name:       item.Coffee.Name,
				Ingredient: ingredients,
			},
			Quantity: item.Quantity,
		}

		orderItems = append(orderItems, oi)
	}

	order, err := client.CreateCustomOrder(orderItems)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	ui.Say(fmt.Sprintf("Order %d created!", order.ID))

	// Set the value of the generated data that will become available to provisioners.
	state.Put("generated_data", map[string]interface{}{
		"OrderId": strconv.Itoa(order.ID),
	})

	state.Put("order", order)
	return multistep.ActionContinue
}

func (s *StepCreateOrder) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
