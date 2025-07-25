
data "datapower_conformancepolicy" "test" {
  depends_on = [datapower_conformancepolicy.test]
  app_domain = "acc_test_domain"
}