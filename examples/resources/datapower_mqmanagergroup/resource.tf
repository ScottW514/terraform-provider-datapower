
resource "datapower_mqmanagergroup" "test" {
  id                    = "ResTestMQManagerGroup"
  app_domain            = "acceptance_test"
  primary_queue_manager = "AccTest_MQManager"
}