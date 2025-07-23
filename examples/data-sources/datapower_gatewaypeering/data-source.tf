
data "datapower_gatewaypeering" "test" {
  depends_on = [datapower_gatewaypeering.test]
  app_domain = "acc_test_domain"
}