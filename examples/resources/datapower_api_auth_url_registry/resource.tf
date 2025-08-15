
resource "datapower_api_auth_url_registry" "test" {
  id         = "ResTestAPIAuthURLRegistry"
  app_domain = "acceptance_test"
  auth_url   = "http://localhost"
}