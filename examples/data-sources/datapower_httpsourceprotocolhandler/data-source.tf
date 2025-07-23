
data "datapower_httpsourceprotocolhandler" "test" {
  depends_on = [datapower_httpsourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}