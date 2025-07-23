
resource "datapower_apiplan" "test" {
  id               = "APIPlan_test"
  app_domain       = "acc_test_domain"
  api              = [datapower_apidefinition.test.id]
  rate_limit_scope = "per-application"
}