
data "datapower_wsstylepolicyrule" "test" {
  depends_on = [datapower_wsstylepolicyrule.test]
  app_domain = "acc_test_domain"
}