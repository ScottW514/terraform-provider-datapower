
data "datapower_apidefinition" "test" {
  depends_on = [datapower_apidefinition.test]
  app_domain = "acc_test_domain"
}