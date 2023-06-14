The HashiCups plugin is part of the [Packer](https://learn.hashicorp.com/packer) Learn collection.
The plugin interacts with the [HashiCorp Demo App API](https://github.com/hashicorp-demoapp/product-api-go) called HashiCups.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    hashicups = {
      source  = "github.com/hashicorp/hashicups"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/hashicups
```

### Components

#### Builders

- [Order builder](/packer/integrations/hashicorp/hashicups/latest/components/builder/order) - The order builder is used to create custom HashiCups order.

#### Provisioners

- [Toppings provisioner](/packer/integrations/hashicorp/hashicups/latest/components/provisioner/toppings) - The toppings provisioner is used to add toppings to your order coffee(s).

#### Post-processors

- [Receipt post-processor](/packer/integrations/hashicorp/hashicups/latest/components/post-processor/receipt) - The receipt post-processor is used to
  print the receipt of your order.

#### Data Sources

- [Coffees data source](/packer/integrations/hashicorp/hashicups/latest/components/data-source/coffees) - The coffees data source is used to
  fetch all the coffees ids existent in the HashiCups menu.

- [Ingredients data source](/packer/integrations/hashicorp/hashicups/latest/components/data-source/ingredients) - The ingredients data source is used to
  fetch the ingredients ids for an existent coffee in the HashiCups menu.

### The HashiCups menu and orders

Get the available coffees:
```shell
$ curl -v localhost:19090/coffees | jq
```

The following api call requires authorization.

First, sign-in with previously created account:
```shell
$ curl -X POST localhost:19090/signin -d '{"username":"education", "password":"test123"}'
```

Then, export the returned JWT token to `HASHICUPS_TOKEN`:
```shell
$ export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU5ODcxNzgsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiZWR1Y2F0aW9uIn0.VJQXoxror-_ZUoNHtsG6GJ-bJCOvjU5kMZtXhSzBtP0
```
With that, you can perform authorized calls.

Get the ingredients for a coffee:

````shell
$ curl -X GET  -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/coffees/1/ingredients | jq
````

Get the created orders:
```shell
$ curl -X GET  -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders | jq
```
