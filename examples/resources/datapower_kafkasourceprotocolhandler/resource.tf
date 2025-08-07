
resource "datapower_kafkasourceprotocolhandler" "test" {
  id             = "ResTestKafkaSourceProtocolHandler"
  app_domain     = "acceptance_test"
  cluster        = "AccTest_KafkaCluster"
  request_topic  = "topic"
  consumer_group = "consumer"
}