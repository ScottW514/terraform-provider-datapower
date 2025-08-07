
resource "datapower_durationmonitor" "test" {
  id           = "ResTestDurationMonitor"
  app_domain   = "acceptance_test"
  measure      = "messages"
  message_type = "AccTest_MessageType"
}