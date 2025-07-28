
resource "datapower_countmonitor" "test" {
  id           = "CookieAttributePolicy_name"
  app_domain   = "acc_test_domain"
  measure      = "requests"
  source       = "all"
  max_sources  = 10000
  message_type = "TestAccMessageType"
}