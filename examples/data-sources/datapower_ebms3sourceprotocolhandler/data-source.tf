
data "datapower_ebms3sourceprotocolhandler" "test" {
  depends_on = [datapower_ebms3sourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}