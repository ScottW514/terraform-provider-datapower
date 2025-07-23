
resource "datapower_apiauthurlregistry" "test" {
  id         = "APIAuthURLRegistry_test"
  app_domain = "acc_test_domain"
  auth_url   = "http://localhost"
}