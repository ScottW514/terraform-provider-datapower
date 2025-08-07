
resource "datapower_mcfhttpheader" "test" {
  id         = "ResTestMCFHttpHeader"
  app_domain = "acceptance_test"
  http_name  = "HEADERNAME"
  http_value = "HEADERVALUE"
}