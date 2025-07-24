
data "datapower_mqmanager" "test" {
  depends_on = [datapower_mqmanager.test]
  app_domain = "acc_test_domain"
}