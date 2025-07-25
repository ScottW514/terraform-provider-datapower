
data "datapower_wsstylepolicy" "test" {
  depends_on = [datapower_wsstylepolicy.test]
  app_domain = "acc_test_domain"
}