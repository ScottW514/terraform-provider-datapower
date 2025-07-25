
resource "datapower_kafkacluster" "test" {
  id         = "KafkaCluster_name"
  app_domain = "acc_test_domain"
  protocol   = "plaintext"
  endpoint = [{
    host = "localhost"
    port = 8888
  }]
}