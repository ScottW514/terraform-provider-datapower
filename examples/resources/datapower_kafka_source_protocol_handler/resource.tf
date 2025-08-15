
resource "datapower_kafka_source_protocol_handler" "test" {
  id             = "ResTestKafkaSourceProtocolHandler"
  app_domain     = "acceptance_test"
  cluster        = "AccTest_KafkaCluster"
  request_topic  = "topic"
  consumer_group = "consumer"
}