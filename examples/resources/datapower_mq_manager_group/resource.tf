
resource "datapower_mq_manager_group" "test" {
  id                    = "ResTestMQManagerGroup"
  app_domain            = "acceptance_test"
  primary_queue_manager = "AccTest_MQManager"
}