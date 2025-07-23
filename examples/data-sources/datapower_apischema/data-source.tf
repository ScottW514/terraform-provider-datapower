
data "datapower_apischema" "test" {
  depends_on = [datapower_apischema.test]
  app_domain = "acc_test_domain"
}