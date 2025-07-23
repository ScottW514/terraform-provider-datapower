
data "datapower_apigateway" "test" {
  depends_on = [datapower_apigateway.test]
  app_domain = "acc_test_domain"
}