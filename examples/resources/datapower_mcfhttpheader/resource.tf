
resource "datapower_mcfhttpheader" "test" {
  id         = "MCFHttpHeader_name"
  app_domain = "acc_test_domain"
  http_name  = "HEADERNAME"
  http_value = "HEADERVALUE"
}