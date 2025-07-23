
data "datapower_socialloginpolicy" "test" {
  depends_on = [datapower_socialloginpolicy.test]
  app_domain = "acc_test_domain"
}