package order

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepCreateOrder struct {
	Items []OrderItem
}

func (s *StepCreateOrder) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	client := state.Get("client").(*hashicups.Client)

	ui.Say("Creating HashiCups Order")

	originalCoffees, err := client.GetCoffees()
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	orderItems := []hashicups.OrderItem{}
	for _, item := range s.Items {
		customIngredients := item.Coffee.Ingredient
		orderItem := hashicups.OrderItem{
			Coffee: hashicups.Coffee{
				Name: item.Coffee.Name,
			},
			Quantity: item.Quantity,
		}

		originalIngredients, err := client.GetCoffeeIngredients(item.Coffee.ID)
		if err != nil {
			state.Put("error", fmt.Errorf("pelase, pick a valid coffee id from the menu: %s", err.Error()))
			return multistep.ActionHalt
		}

		// Update order with newly create custom coffee
		for _, coffee := range originalCoffees {
			if strconv.Itoa(coffee.ID) == item.Coffee.ID {
				if orderItem.Coffee.Name == coffee.Name {
					state.Put("error", fmt.Errorf("coffee %s must have a different name from the original coffee", orderItem.Coffee.Name))
					return multistep.ActionHalt
				}
				coffee.Name = orderItem.Coffee.Name
				newCoffee, err := client.CreateCoffee(coffee)
				if err != nil {
					state.Put("error", err)
					return multistep.ActionHalt
				}
				orderItem.Coffee = *newCoffee
				continue
			}
		}

		// Add ingredients to the custom coffee
		for _, ingredient := range originalIngredients {
			for _, customIngredient := range customIngredients {
				if strconv.Itoa(ingredient.ID) == customIngredient.ID {
					// Update ingredient quantity according to customisation
					ingredient.Quantity = customIngredient.Quantity
					continue
				}
			}
			_, err := client.CreateCoffeeIngredient(orderItem.Coffee, ingredient)
			if err != nil {
				state.Put("error", err)
				return multistep.ActionHalt
			}
		}

		orderItems = append(orderItems, orderItem)
	}

	order, err := client.CreateOrder(orderItems)
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
