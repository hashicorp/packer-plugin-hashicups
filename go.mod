module packer-plugin-hashicups

go 1.16

require (
	github.com/hashicorp-demoapp/hashicups-client-go v0.0.0-20210316102605-c2dc0b667c4e
	github.com/hashicorp/hcl/v2 v2.16.2
	github.com/hashicorp/packer-plugin-sdk v0.5.0
	github.com/jung-kurt/gofpdf v1.16.2
	github.com/mitchellh/mapstructure v1.4.1
	github.com/zclconf/go-cty v1.12.1
)

replace github.com/zclconf/go-cty => github.com/nywilken/go-cty v1.12.1 // added by packer-sdc fix as noted in github.com/hashicorp/packer-plugin-sdk/issues/187
