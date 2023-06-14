Type: `hashicups-ingredients`

The ingredients data source is used to fetch the ingredients ids for an existent coffee in the HashiCups menu.

## Required

<!-- Code generated from the comments of the AuthConfig struct in common/auth.go; DO NOT EDIT MANUALLY -->

- `username` (string) - The username signed up to the Product API.

- `password` (string) - The password for the username signed up to the Product API.

<!-- End of code generated from the comments of the AuthConfig struct in common/auth.go; -->


- `coffee` (string) - The coffee id you would like to get the ingredient from. The ID should exist in the HashiCups menu.

## Optional

<!-- Code generated from the comments of the AuthConfig struct in common/auth.go; DO NOT EDIT MANUALLY -->

- `host` (string) - The Product API host URL. Defaults to `localhost:19090`

<!-- End of code generated from the comments of the AuthConfig struct in common/auth.go; -->


## OutPut

- `map` (map[string]string) - A map of ingredient name to ingredient id.

## Example Usage

```hcl
data "hashicups-ingredients" "packer-ingredients" {
  username = "education"
  password = "test123"
  coffee = "Packer Spiced Latte"
}

locals {
  semi_skimmed_milk = data.hashicups-ingredients.packer-ingredients.map["Semi Skimmed Milk"]
}

source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

 item {
    coffee {
      id = 1
      name = "my custom packer spiced latter"
      ingredient {
        id = local.semi_skimmed_milk
        quantity = 200
      }
    }
    quantity = 2
  }
}

build {
  sources = ["sources.hashicups-order.my-custom-order"]
}
```
