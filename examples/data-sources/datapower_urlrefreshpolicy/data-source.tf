
data "datapower_urlrefreshpolicy" "test" {
  depends_on = [datapower_urlrefreshpolicy.test]
  app_domain = "acc_test_domain"
}