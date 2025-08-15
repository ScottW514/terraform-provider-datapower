
resource "datapower_api_schema" "test" {
  id          = "ResTestAPISchema"
  app_domain  = "acceptance_test"
  json_schema = "http://localhost/json"
}