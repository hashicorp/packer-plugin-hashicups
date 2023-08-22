Type: `hashicups-receipt`

The receipt post-processor is used to print the receipt of your order.

## Optional

- `filename` (string) - The receipt file name. Should not include the extension. Defaults to `receipt`.
- `format` (string) - The receipt format. Should be `pdf` or `txt`. Defaults to `pdf`.

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

  post-processor "hashicups-receipt" {
    format = "pdf"
  }
}
```
