
resource "datapower_duration_monitor" "test" {
  id         = "ResTestDurationMonitor"
  app_domain = "acceptance_test"
  measure    = "messages"
  filter = [{
    name   = "ResTestDmDurationMonitorFilter"
    value  = 1
    action = "AccTest_FilterAction"
  }]
  message_type = "AccTest_MessageType"
}