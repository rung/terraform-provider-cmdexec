package cmdexec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"cmdexec_output": ResourceCmdOutput(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cmdexec_execute": dataSourceCmdExecExecute(),
		},
	}
	return provider
}
