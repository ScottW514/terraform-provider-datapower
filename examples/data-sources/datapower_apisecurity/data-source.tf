
data "datapower_apisecurity" "test" {
  depends_on = [datapower_apisecurity.test]
  app_domain = "acc_test_domain"
}