
resource "datapower_apisecuritybasicauth" "test" {
  id            = "APISecurityBasicAuth_name"
  app_domain    = "acc_test_domain"
  user_registry = datapower_apiauthurlregistry.test.id
}