
resource "datapower_amqpsourceprotocolhandler" "test" {
  id         = "AMQPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  broker     = datapower_amqpbroker.test.id
  from       = "amqpfrom"
}