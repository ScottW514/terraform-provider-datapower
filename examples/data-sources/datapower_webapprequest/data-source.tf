
data "datapower_webapprequest" "test" {
  depends_on = [datapower_webapprequest.test]
  app_domain = "acc_test_domain"
}