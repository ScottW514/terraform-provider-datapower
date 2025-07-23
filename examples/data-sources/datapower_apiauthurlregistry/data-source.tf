
data "datapower_apiauthurlregistry" "test" {
  depends_on = [datapower_apiauthurlregistry.test]
  app_domain = "acc_test_domain"
}