
data "datapower_wxsgrid" "test" {
  depends_on = [datapower_wxsgrid.test]
  app_domain = "acc_test_domain"
}