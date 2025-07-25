
data "datapower_xtcprotocolhandler" "test" {
  depends_on = [datapower_xtcprotocolhandler.test]
  app_domain = "acc_test_domain"
}