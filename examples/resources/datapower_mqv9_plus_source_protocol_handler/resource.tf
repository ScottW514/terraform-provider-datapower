
resource "datapower_mqv9_plus_source_protocol_handler" "test" {
  id            = "ResTestMQv9PlusSourceProtocolHandler"
  app_domain    = "acceptance_test"
  queue_manager = "AccTest_MQManager"
}