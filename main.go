package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/rung/terraform-provider-cmdexec/cmdexec"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cmdexec.Provider})
}
