
data "datapower_webservicesagent" "test" {
  depends_on = [datapower_webservicesagent.test]
  app_domain = "acc_test_domain"
}