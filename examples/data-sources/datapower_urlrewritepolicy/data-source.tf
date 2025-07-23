
data "datapower_urlrewritepolicy" "test" {
  depends_on = [datapower_urlrewritepolicy.test]
  app_domain = "acc_test_domain"
}