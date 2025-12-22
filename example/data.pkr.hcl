# Copyright IBM Corp. 2021, 2025
# SPDX-License-Identifier: MPL-2.0

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