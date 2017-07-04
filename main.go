package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/idmsubs/terraform-provider-odl/odl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: odl.Provider})
}
