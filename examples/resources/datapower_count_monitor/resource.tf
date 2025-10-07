
resource "datapower_count_monitor" "test" {
  id         = "ResTestCountMonitor"
  app_domain = "acceptance_test"
  measure    = "requests"
  source     = "all"
  filter = [{
    name        = "Filter1"
    interval    = 1000
    rate_limit  = 50
    burst_limit = 100
    action      = "AccTest_FilterAction"
  }]
  max_source_s = 10000
  message_type = "AccTest_MessageType"
}