
resource "datapower_count_monitor" "test" {
  id           = "ResTestCountMonitor"
  app_domain   = "acceptance_test"
  measure      = "requests"
  source       = "all"
  max_source_s = 10000
  message_type = "AccTest_MessageType"
}