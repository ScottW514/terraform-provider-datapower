
data "datapower_webservicemonitor" "test" {
  depends_on = [datapower_webservicemonitor.test]
  app_domain = "acc_test_domain"
}