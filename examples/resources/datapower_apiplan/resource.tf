
resource "datapower_apiplan" "test" {
  id               = "APIPlan_name"
  app_domain       = "acc_test_domain"
  api              = [datapower_apidefinition.test.id]
  rate_limit_scope = "per-application"
}