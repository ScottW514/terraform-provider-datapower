
resource "datapower_mcfhttpurl" "test" {
  id                  = "ResTestMCFHttpURL"
  app_domain          = "acceptance_test"
  http_url_expression = "*"
}