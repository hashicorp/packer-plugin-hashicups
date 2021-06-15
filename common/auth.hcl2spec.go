// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatAuthConfig is an auto-generated flat version of AuthConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatAuthConfig struct {
	Username *string `mapstructure:"username" cty:"username" hcl:"username"`
	Password *string `mapstructure:"password" cty:"password" hcl:"password"`
	Host     *string `mapstructure:"host" cty:"host" hcl:"host"`
}

// FlatMapstructure returns a new FlatAuthConfig.
// FlatAuthConfig is an auto-generated flat version of AuthConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*AuthConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatAuthConfig)
}

// HCL2Spec returns the hcl spec of a AuthConfig.
// This spec is used by HCL to read the fields of AuthConfig.
// The decoded values from this spec will then be applied to a FlatAuthConfig.
func (*FlatAuthConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"username": &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password": &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"host":     &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
	}
	return s
}
