
resource "datapower_mqv9plussourceprotocolhandler" "test" {
  id            = "ResTestMQv9PlusSourceProtocolHandler"
  app_domain    = "acceptance_test"
  queue_manager = "AccTest_MQManager"
}