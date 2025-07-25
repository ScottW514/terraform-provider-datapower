
data "datapower_controllist" "test" {
  depends_on = [datapower_controllist.test]
  app_domain = "acc_test_domain"
}