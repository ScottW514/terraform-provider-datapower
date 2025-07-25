
resource "datapower_mcfhttpmethod" "test" {
  id          = "MCFHttpMethod_name"
  app_domain  = "acc_test_domain"
  http_method = "GET"
}