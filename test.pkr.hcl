data "hashicups-coffees" "coffees" {
  username = "education"
  password = "test123"
}

data "hashicups-ingredients" "ingredients" {
  username = "education"
  password = "test123"
  coffee = "Vagrante espresso"
}

locals {
  vagrante_espresso = data.hashicups-coffees.coffees.map["Vagrante espresso"]
  espresso = data.hashicups-ingredients.ingredients.map["Espresso"]
}

source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

  item {
    coffee {
      id = local.vagrante_espresso
      name = "my new custom coffee"
      ingredient {
        id = local.espresso
        quantity = 50
      }
    }
    quantity = 1
  }
}

build {
  sources = ["sources.hashicups-order.my-custom-order"]

  provisioner "hasicups-status" {
    order = build.OrderId
  }

//  post-processor "hashicups-recipt" {
//    order = build.OrderId
//    format = "pdf"
//  }
}