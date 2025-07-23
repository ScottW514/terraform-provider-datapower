
data "datapower_xmlmanager" "test" {
  depends_on = [datapower_xmlmanager.test]
  app_domain = "acc_test_domain"
}