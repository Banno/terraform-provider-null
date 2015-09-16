package main

import (
	"github.com/banno/terraform-provider-null/null"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: null.Provider,
	})
}
