
resource "datapower_log_target" "test" {
  id         = "ResTest_LogTarget"
  app_domain = "acceptance_test"
  type       = "file"
  local_file = "logtemp:///AccTest_LogTarget.log"
}