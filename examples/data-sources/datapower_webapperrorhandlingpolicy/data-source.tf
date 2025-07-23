
data "datapower_webapperrorhandlingpolicy" "test" {
  depends_on = [datapower_webapperrorhandlingpolicy.test]
  app_domain = "acc_test_domain"
}