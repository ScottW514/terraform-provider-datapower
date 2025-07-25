
data "datapower_mtompolicy" "test" {
  depends_on = [datapower_mtompolicy.test]
  app_domain = "acc_test_domain"
}