
data "datapower_aaapolicy" "test" {
  depends_on = [datapower_aaapolicy.test]
  app_domain = "acc_test_domain"
}