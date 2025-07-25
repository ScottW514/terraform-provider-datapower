
data "datapower_gitopstemplate" "test" {
  depends_on = [datapower_gitopstemplate.test]
  app_domain = "acc_test_domain"
}