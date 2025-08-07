
resource "datapower_mqv9plusmftsourceprotocolhandler" "test" {
  id            = "ResTestMQv9PlusMFTSourceProtocolHandler"
  app_domain    = "acceptance_test"
  queue_manager = "AccTest_MQManager"
  get_queue     = "queue"
}