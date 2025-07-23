
data "datapower_apiplan" "test" {
  depends_on = [datapower_apiplan.test]
  app_domain = "acc_test_domain"
}