
data "datapower_cookieattributepolicy" "test" {
  depends_on = [datapower_cookieattributepolicy.test]
  app_domain = "acc_test_domain"
}