package cmdexec

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"os/exec"
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
	sh, err := exec.LookPath("sh")
	if err != nil {
		return errors.New("sh is not found")
	}

	command := d.Get("command").(string)
	timeout := d.Get("timeout").(int)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	cmd := exec.CommandContext(ctx, sh, "-c", command)
	output, _ := cmd.CombinedOutput()

	d.Set("rc", cmd.ProcessState.ExitCode())
	d.Set("output", string(output))
	d.SetId("0")

	return nil
}
