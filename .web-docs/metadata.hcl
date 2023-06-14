# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "HashiCups"
  description = "The HashiCups plugin is a reference plugin for learning how to use HashiCorp Packer."
  identifier = "packer/hashicorp/hashicups"
  component {
    type = "data-source"
    name = "Coffees"
    slug = "coffees"
  }
  component {
    type = "data-source"
    name = "Ingredients"
    slug = "ingredients"
  }
  component {
    type = "builder"
    name = "Order"
    slug = "order"
  }
  component {
    type = "provisioner"
    name = "Toppings"
    slug = "toppings"
  }
  component {
    type = "post-processor"
    name = "Receipt"
    slug = "receipt"
  }
}
