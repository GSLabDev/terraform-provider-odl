package main

import (
	"github.com/GSLabDev/terraform-provider-odl/odl"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: odl.Provider})
}
