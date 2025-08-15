
resource "datapower_amqp_source_protocol_handler" "test" {
  id         = "ResTestAMQPSourceProtocolHandler"
  app_domain = "acceptance_test"
  broker     = "AccTest_AMQPBroker"
  from       = "amqpfrom"
}