
data "datapower_as3sourceprotocolhandler" "test" {
  depends_on = [datapower_as3sourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}