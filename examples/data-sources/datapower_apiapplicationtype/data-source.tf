
data "datapower_apiapplicationtype" "test" {
  depends_on = [datapower_apiapplicationtype.test]
  app_domain = "acc_test_domain"
}