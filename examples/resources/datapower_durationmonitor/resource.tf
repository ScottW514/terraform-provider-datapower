
resource "datapower_durationmonitor" "test" {
  id           = "DurationMonitor_name"
  app_domain   = "acc_test_domain"
  measure      = "messages"
  message_type = "TestAccMessageType"
}