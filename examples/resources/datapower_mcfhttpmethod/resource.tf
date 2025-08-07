
resource "datapower_mcfhttpmethod" "test" {
  id          = "ResTestMCFHttpMethod"
  app_domain  = "acceptance_test"
  http_method = "GET"
}