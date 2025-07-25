
data "datapower_corsrule" "test" {
  depends_on = [datapower_corsrule.test]
  app_domain = "acc_test_domain"
}