
resource "datapower_kafkasourceprotocolhandler" "test" {
  id             = "KafkaSourceProtocolHandler_name"
  app_domain     = "acc_test_domain"
  cluster        = datapower_kafkacluster.test.id
  request_topic  = "topic"
  consumer_group = "consumer"
}