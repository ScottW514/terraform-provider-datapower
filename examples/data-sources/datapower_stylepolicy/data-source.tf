
data "datapower_stylepolicy" "test" {
  depends_on = [datapower_stylepolicy.test]
  app_domain = "acc_test_domain"
}