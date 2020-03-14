package cmdexec

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"os/exec"
	"strings"
	"time"
)

func dataSourceCmdExecExecute() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCmdExecExecuteRead,

		Schema: map[string]*schema.Schema{
			"command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"rc": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"output": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCmdExecExecuteRead(d *schema.ResourceData, meta interface{}) error {
	cmd := d.Get("command").(string)
	timeout := d.Get("timeout").(int)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	out, err := exec.CommandContext(ctx, strings.Split(cmd, " ")[0], strings.Split(cmd, " ")[1:]...).Output()
	if err != nil {
		return errors.New("it failed to execute command")
	}

	d.Set("output", string(out))
	return nil
}
