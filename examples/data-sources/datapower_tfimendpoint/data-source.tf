
data "datapower_tfimendpoint" "test" {
  depends_on = [datapower_tfimendpoint.test]
  app_domain = "acc_test_domain"
}