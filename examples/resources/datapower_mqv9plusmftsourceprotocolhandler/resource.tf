
resource "datapower_mqv9plusmftsourceprotocolhandler" "test" {
  id            = "MQv9PlusMFTSourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  queue_manager = datapower_mqmanager.test.id
  get_queue     = "queue"
}