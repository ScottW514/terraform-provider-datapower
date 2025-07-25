
data "datapower_tcpproxyservice" "test" {
  depends_on = [datapower_tcpproxyservice.test]
  app_domain = "acc_test_domain"
}