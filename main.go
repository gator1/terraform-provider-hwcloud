package main

import (
	"github.com/gator1/terraform-provider-hwcloud/huaweicloud" // TODO: Revert path when merge
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: huaweicloud.Provider})
}
