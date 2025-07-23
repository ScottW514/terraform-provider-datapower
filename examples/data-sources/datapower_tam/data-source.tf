
data "datapower_tam" "test" {
  depends_on = [datapower_tam.test]
  app_domain = "acc_test_domain"
}