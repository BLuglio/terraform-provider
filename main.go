package main

import (
	"github.com/BLuglio/terraform-provider/custom_provider"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: custom_provider.Provider,
	})
}
