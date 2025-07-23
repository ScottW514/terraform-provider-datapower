
data "datapower_apiresult" "test" {
  depends_on = [datapower_apiresult.test]
  app_domain = "acc_test_domain"
}