
resource "datapower_analytics_endpoint" "test" {
  id                     = "ResTestAnalyticsEndpoint"
  app_domain             = "acceptance_test"
  analytics_server_url   = "https://localhost"
  ssl_client             = "AccTest_SSLClientProfile"
  max_records            = 1024
  max_records_memory_kb  = 512
  max_delivery_memory_mb = 512
}