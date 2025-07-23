
data "datapower_webappsessionpolicy" "test" {
  depends_on = [datapower_webappsessionpolicy.test]
  app_domain = "acc_test_domain"
}