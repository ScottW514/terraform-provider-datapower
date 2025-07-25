
data "datapower_corspolicy" "test" {
  depends_on = [datapower_corspolicy.test]
  app_domain = "acc_test_domain"
}