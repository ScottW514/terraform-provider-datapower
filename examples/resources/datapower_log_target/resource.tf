
resource "datapower_log_target" "test" {
  id         = "ResTest_LogTarget"
  app_domain = "acceptance_test"
  log_events = [{
    class = "AccTest_LogLabel"
  }]
  type                       = "file"
  local_file                 = "logtemp:///AccTest_LogTarget.log"
  remote_password_wo_version = 1
  log_objects = [{
    class = "AAAPolicy"
  }]
  log_ip_filter = null
  log_triggers = [{
  }]
}