
resource "datapower_mcf_http_header" "test" {
  id         = "ResTestMCFHttpHeader"
  app_domain = "acceptance_test"
  http_name  = "HEADERNAME"
  http_value = "HEADERVALUE"
}