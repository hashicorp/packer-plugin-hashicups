Type: `hashicups-toppings`

The toppings provisioner is used to add toppings to your order coffee(s).

## Required

- `toppings` ([]string) - Any toppings you would like to pour on top of the ordered coffees.

## Example Usage


```hcl
source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

  item {
    coffee {
      id = 1
      name = "My Custom Packer Spiced Latter"
      ingredient {
        id = 1
        quantity = 50
      }
    }
  }
}

build {
  sources = ["sources.hashicups-order.my-custom-order"]

  provisioner "hashicups-toppings" {
    toppings = ["cinnamon", "marshmallow", "chocolate", "sprinkles"]
  }
}
```
