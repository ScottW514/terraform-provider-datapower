
data "datapower_amqpbroker" "test" {
  depends_on = [datapower_amqpbroker.test]
  app_domain = "acc_test_domain"
}