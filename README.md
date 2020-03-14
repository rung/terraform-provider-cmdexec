# terraform-provider-cmdexec

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html)
- [Go](https://golang.org/doc/install) (to build the provider plugin)


Examples
---------------------

To install the provider, please run `make install`.
```sh
make install
```

To use the provider, please write datasource.
value of `command` is executed.
```tf
data "cmdexec_execute" "sample" {
  command = "echo test"
}
```

To get output of the executed command when `terraform plan`, please write resource.
```tf
resource "cmdexec_output" "sample" {
  rc     = data.cmdexec_execute.sample.rc
  output = data.cmdexec_execute.sample.output
}
```
