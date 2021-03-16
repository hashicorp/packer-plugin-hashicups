//go:generate mapstructure-to-hcl2 -type AuthConfig

package common

import "github.com/hashicorp-demoapp/hashicups-client-go"

type AuthConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
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
