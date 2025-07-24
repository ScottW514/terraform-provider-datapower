
resource "datapower_apiauthurlregistry" "test" {
  id         = "APIAuthURLRegistry_name"
  app_domain = "acc_test_domain"
  auth_url   = "http://localhost"
}