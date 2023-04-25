// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type AuthConfig

package common

import "github.com/hashicorp-demoapp/hashicups-client-go"

type AuthConfig struct {
	// The username signed up to the Product API.
	Username string `mapstructure:"username" required:"true"`
	// The password for the username signed up to the Product API.
	Password string `mapstructure:"password" required:"true"`
	// The Product API host URL. Defaults to `localhost:19090`
	Host string `mapstructure:"host"`
}

func (c *AuthConfig) CreateClient() (*hashicups.Client, error) {
	var host *string
	var username *string
	var password *string

	if c.Host != "" {
		host = &c.Host
	}
	if c.Username != "" {
		username = &c.Username
	}
	if c.Password != "" {
		password = &c.Password
	}
	return hashicups.NewClient(host, username, password)
}
