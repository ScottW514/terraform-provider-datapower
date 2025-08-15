
resource "datapower_duration_monitor" "test" {
  id           = "ResTestDurationMonitor"
  app_domain   = "acceptance_test"
  measure      = "messages"
  message_type = "AccTest_MessageType"
}