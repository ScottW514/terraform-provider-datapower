
resource "datapower_apiauthurlregistry" "test" {
  id         = "ResTestAPIAuthURLRegistry"
  app_domain = "acceptance_test"
  auth_url   = "http://localhost"
}