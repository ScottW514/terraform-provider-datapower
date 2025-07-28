
resource "datapower_amqpsourceprotocolhandler" "test" {
  id         = "AMQPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  broker     = "TestAccAMQPBroker"
  from       = "amqpfrom"
}