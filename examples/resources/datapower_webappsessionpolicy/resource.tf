
resource "datapower_webappsessionpolicy" "test" {
  id            = "WebAppSessionPolicy_name"
  app_domain    = "acc_test_domain"
  start_matches = "__default-accept-service-providers__"
}