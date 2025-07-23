
data "datapower_webappresponse" "test" {
  depends_on = [datapower_webappresponse.test]
  app_domain = "acc_test_domain"
}