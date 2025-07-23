
data "datapower_stylepolicyrule" "test" {
  depends_on = [datapower_stylepolicyrule.test]
  app_domain = "acc_test_domain"
}