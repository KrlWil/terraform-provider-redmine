package main

import (
	"github.com/KrlWil/terraform-provider-redmine/redmine"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return redmine.Provider()
		},

	})
}
