package redmine

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider: returns a terraform.ResourceProvider and maps redmine_project to it
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"redmine_project": resourceProject(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure configures the provider by creating and authenticating Redmine client from config.go
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var c Config

    c.createAndAuthenticateClient()
	return &c, nil
}

