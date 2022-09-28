package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	cc "go.clever-cloud.dev/steampipe/clevercloud"
)

func main() {
	opts := &plugin.ServeOpts{
		PluginName: "steampipe-plugin-clevercloud",
		PluginFunc: cc.Plugin,
	}

	plugin.Serve(opts)
}
