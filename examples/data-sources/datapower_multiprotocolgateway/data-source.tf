
data "datapower_multiprotocolgateway" "test" {
  depends_on = [datapower_multiprotocolgateway.test]
  app_domain = "acc_test_domain"
}