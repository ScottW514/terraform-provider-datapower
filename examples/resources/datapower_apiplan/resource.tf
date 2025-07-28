
resource "datapower_apiplan" "test" {
  id               = "APIPlan_name"
  app_domain       = "acc_test_domain"
  api              = ["TestAccAPIDefinition"]
  rate_limit_scope = "per-application"
}