
data "datapower_appsecuritypolicy" "test" {
  depends_on = [datapower_appsecuritypolicy.test]
  app_domain = "acc_test_domain"
}