
data "datapower_as2sourceprotocolhandler" "test" {
  depends_on = [datapower_as2sourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}