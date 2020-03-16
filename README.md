# terraform-provider-cmdexec

What is this
------------

`terraform-provider-cmdexec` provides command execution from Terraform Configuration.  
Terraform has `local-exec` provisioner by default. but provisioner is executed when `terraform apply`. On the other hand, `terraform-provider-cmdexec` execute a command when `terraform plan`.  
This provider was originally created for penetration testing of CI/CD pipeline.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html)
- `/bin/sh`
- [Go](https://golang.org/doc/install) (to build the provider plugin)


How to use
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

To build the provider for Linux(x64), please run `make build-linux`
```sh
make build-linux
```
