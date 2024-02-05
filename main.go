package main

import (
	"flag"

	provider "github.com/Nmishin/terraform-provider-dodo/internal/sdkv2provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	version string = "dev"
	commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider(version, false)
		},
		Debug: debugMode,
	}

	plugin.Serve(opts)
}
