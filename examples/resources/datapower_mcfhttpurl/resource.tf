
resource "datapower_mcfhttpurl" "test" {
  id                  = "MCFHttpURL_name"
  app_domain          = "acc_test_domain"
  http_url_expression = "*"
}