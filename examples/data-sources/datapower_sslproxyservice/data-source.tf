
data "datapower_sslproxyservice" "test" {
  depends_on = [datapower_sslproxyservice.test]
  app_domain = "acc_test_domain"
}