
data "datapower_mqv9plussourceprotocolhandler" "test" {
  depends_on = [datapower_mqv9plussourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}