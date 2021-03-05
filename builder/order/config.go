//go:generate mapstructure-to-hcl2 -type Config,OrderItem,Coffee,Ingredient

package order

import (
	packercommon "github.com/hashicorp/packer-plugin-sdk/common"
	"packer-plugin-hashicups/common"
)

type Config struct {
	packercommon.PackerConfig `mapstructure:",squash"`
	common.AuthConfig         `mapstructure:",squash"`
	Item                      []OrderItem `mapstructure:"item,omitempty" required:"true"`
}

type OrderItem struct {
	Coffee   Coffee `mapstructure:"coffee" required:"true"`
	Quantity int    `mapstructure:"quantity" required:"true"`
}

type Coffee struct {
	ID         string       `mapstructure:"id" required:"true"`
	Name       string       `mapstructure:"name" required:"true"`
	Ingredient []Ingredient `mapstructure:"ingredient"`
}

type Ingredient struct {
	ID       string `mapstructure:"id" required:"true"`
	Quantity int    `mapstructure:"quantity"`
}
