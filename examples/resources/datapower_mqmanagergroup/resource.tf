
resource "datapower_mqmanagergroup" "test" {
  id                    = "MQManagerGroup_name"
  app_domain            = "acc_test_domain"
  primary_queue_manager = "TestAccMQManager"
}