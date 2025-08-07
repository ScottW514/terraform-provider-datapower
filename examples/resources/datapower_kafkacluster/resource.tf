
resource "datapower_kafkacluster" "test" {
  id         = "ResTestKafkaCluster"
  app_domain = "acceptance_test"
  protocol   = "plaintext"
  endpoint = [{
    host = "localhost"
    port = 8888
  }]
}