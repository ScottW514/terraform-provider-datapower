
data "datapower_httpssourceprotocolhandler" "test" {
  depends_on = [datapower_httpssourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}