
resource "datapower_mcf_http_url" "test" {
  id                  = "ResTestMCFHttpURL"
  app_domain          = "acceptance_test"
  http_url_expression = "*"
}