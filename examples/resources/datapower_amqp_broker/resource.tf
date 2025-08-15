
resource "datapower_amqp_broker" "test" {
  id            = "ResTestAMQPBroker"
  app_domain    = "acceptance_test"
  host_name     = "host.name"
  port          = 5672
  xml_manager   = "default"
  authorization = "none"
}