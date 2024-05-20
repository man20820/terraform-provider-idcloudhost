package main

import (
	"github.com/man20820/terraform-provider-idcloudhost/idcloudhost"
	"github.com/hashicorp/terraform-plugin-sdk/tree/main/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/tree/main/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return idcloudhost.Provider()
		},
	})
}
