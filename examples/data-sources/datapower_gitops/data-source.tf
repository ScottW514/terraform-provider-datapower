
data "datapower_gitops" "test" {
  depends_on = [datapower_gitops.test]
  app_domain = "acc_test_domain"
}