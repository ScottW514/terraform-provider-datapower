
data "datapower_as2proxysourceprotocolhandler" "test" {
  depends_on = [datapower_as2proxysourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}