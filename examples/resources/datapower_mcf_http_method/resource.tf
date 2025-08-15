
resource "datapower_mcf_http_method" "test" {
  id          = "ResTestMCFHttpMethod"
  app_domain  = "acceptance_test"
  http_method = "GET"
}