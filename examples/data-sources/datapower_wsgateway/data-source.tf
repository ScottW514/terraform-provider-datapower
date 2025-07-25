
data "datapower_wsgateway" "test" {
  depends_on = [datapower_wsgateway.test]
  app_domain = "acc_test_domain"
}