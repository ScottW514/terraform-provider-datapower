
data "datapower_formsloginpolicy" "test" {
  depends_on = [datapower_formsloginpolicy.test]
  app_domain = "acc_test_domain"
}