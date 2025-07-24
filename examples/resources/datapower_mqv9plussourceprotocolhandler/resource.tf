
resource "datapower_mqv9plussourceprotocolhandler" "test" {
  id            = "MQv9PlusSourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  queue_manager = datapower_mqmanager.test.id
}