
data "datapower_analyticsendpoint" "test" {
  depends_on = [datapower_analyticsendpoint.test]
  app_domain = "acc_test_domain"
}