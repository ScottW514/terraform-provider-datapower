
data "datapower_amqpsourceprotocolhandler" "test" {
  depends_on = [datapower_amqpsourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}