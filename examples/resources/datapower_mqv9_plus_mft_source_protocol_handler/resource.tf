
resource "datapower_mqv9_plus_mft_source_protocol_handler" "test" {
  id            = "ResTestMQv9PlusMFTSourceProtocolHandler"
  app_domain    = "acceptance_test"
  queue_manager = "AccTest_MQManager"
  get_queue     = "queue"
}