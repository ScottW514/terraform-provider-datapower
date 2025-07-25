
data "datapower_lunahagroup" "test" {
  depends_on = [datapower_lunahagroup.test]
  app_domain = "acc_test_domain"
}