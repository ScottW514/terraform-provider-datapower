
data "datapower_apiconnectgatewayservice" "test" {
  depends_on = [datapower_apiconnectgatewayservice.test]
  app_domain = "acc_test_domain"
}