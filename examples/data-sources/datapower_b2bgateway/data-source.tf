
data "datapower_b2bgateway" "test" {
  depends_on = [datapower_b2bgateway.test]
  app_domain = "acc_test_domain"
}