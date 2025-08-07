
resource "datapower_amqpsourceprotocolhandler" "test" {
  id         = "ResTestAMQPSourceProtocolHandler"
  app_domain = "acceptance_test"
  broker     = "AccTest_AMQPBroker"
  from       = "amqpfrom"
}