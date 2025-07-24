
resource "datapower_amqpbroker" "test" {
  id            = "AMQPBroker_name"
  app_domain    = "acc_test_domain"
  host_name     = "host.name"
  port          = 5672
  xml_manager   = "default"
  authorization = "none"
}