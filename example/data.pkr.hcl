data "hashicups-coffees" "coffees" {
  username = "education"
  password = "test123"
}

data "hashicups-ingredients" "vagrante-ingredients" {
  username = "education"
  password = "test123"
  coffee = "Vagrante espresso"
}

data "hashicups-ingredients" "packer-ingredients" {
  username = "education"
  password = "test123"
  coffee = "Packer Spiced Latte"
}