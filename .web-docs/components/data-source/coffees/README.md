Type: `hashicups-coffees`

The coffees data source is used to fetch all the coffees ids existent in the HashiCups menu.

## Required

<!-- Code generated from the comments of the AuthConfig struct in common/auth.go; DO NOT EDIT MANUALLY -->

- `username` (string) - The username signed up to the Product API.

- `password` (string) - The password for the username signed up to the Product API.

<!-- End of code generated from the comments of the AuthConfig struct in common/auth.go; -->


## Optional

<!-- Code generated from the comments of the AuthConfig struct in common/auth.go; DO NOT EDIT MANUALLY -->

- `host` (string) - The Product API host URL. Defaults to `localhost:19090`

<!-- End of code generated from the comments of the AuthConfig struct in common/auth.go; -->


## OutPut

- `map` (map[string]string) - A map of coffee name to coffee id.

## Example Usage


```hcl
data "hashicups-coffees" "coffees" {
  username = "education"
  password = "test123"
}

locals {
  vagrante_espresso = data.hashicups-coffees.coffees.map["Vagrante espresso"]
}

source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

  item {
    coffee {
      id = local.vagrante_espresso
      name = "my custom vagrante"
      ingredient {
        id = 1
        quantity = 50
      }
    }
  }
}

build {
  sources = ["sources.hashicups-order.my-custom-order"]
}
```
