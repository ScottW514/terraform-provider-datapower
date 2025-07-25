
data "datapower_ebms2sourceprotocolhandler" "test" {
  depends_on = [datapower_ebms2sourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}