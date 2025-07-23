
resource "datapower_analyticsendpoint" "test" {
  id                     = "AnalyticsEndpoint_test"
  app_domain             = "acc_test_domain"
  analytics_server_url   = "https://localhost"
  max_records            = 1024
  max_records_memory_kb  = 512
  max_delivery_memory_mb = 512
}