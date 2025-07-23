
data "datapower_compileoptionspolicy" "test" {
  depends_on = [datapower_compileoptionspolicy.test]
  app_domain = "acc_test_domain"
}