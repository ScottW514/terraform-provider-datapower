
data "datapower_apioperation" "test" {
  depends_on = [datapower_apioperation.test]
  app_domain = "acc_test_domain"
}