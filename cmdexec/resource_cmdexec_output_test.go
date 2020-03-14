package cmdexec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccCmdexecOutput(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"cmdexec": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: outputOfEcho(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cmdexec_output.sample", "rc", "0"),
					resource.TestCheckResourceAttr("cmdexec_output.sample", "output", "test\n"),
				),
			},
			{
				Config: outputWithReturnCode(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cmdexec_output.sample", "rc", "15"),
				),
			},
		},
	})
}

func outputOfEcho() string {
	tf := `
data "cmdexec_execute" "sample" {
  command = "echo test"
}

resource "cmdexec_output" "sample" {
  rc     = data.cmdexec_execute.sample.rc
  output = data.cmdexec_execute.sample.output
}
`
	return tf
}

func outputWithReturnCode() string {
	tf := `
data "cmdexec_execute" "sample2" {
  command = "exit 15"
}

resource "cmdexec_output" "sample" {
  rc     = data.cmdexec_execute.sample2.rc
  output = data.cmdexec_execute.sample2.output
}
`
	return tf
}
