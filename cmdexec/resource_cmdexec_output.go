package cmdexec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceCmdOutput() *schema.Resource {
	return &schema.Resource{
		Read: func(*schema.ResourceData, interface{}) error {
			return nil
		},
		Update: func(*schema.ResourceData, interface{}) error {
			return nil
		},
		Create: func(d *schema.ResourceData, i interface{}) error {
			d.SetId("0")
			return nil
		},
		Delete: func(*schema.ResourceData, interface{}) error {
			return nil
		},

		Schema: map[string]*schema.Schema{
			"rc": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  999,
			},
			"output": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
