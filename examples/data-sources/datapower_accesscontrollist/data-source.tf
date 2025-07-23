
data "datapower_accesscontrollist" "test" {
  depends_on = [datapower_accesscontrollist.test]
  app_domain = "acc_test_domain"
}