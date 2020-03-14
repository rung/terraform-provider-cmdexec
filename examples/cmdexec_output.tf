resource "cmdexec_output" "sample" {
  rc     = data.cmdexec_execute.sample.rc
  output = data.cmdexec_execute.sample.output
}
