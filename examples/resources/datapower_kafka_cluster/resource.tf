
resource "datapower_kafka_cluster" "test" {
  id         = "ResTestKafkaCluster"
  app_domain = "acceptance_test"
  protocol   = "plaintext"
  endpoint = [{
    host = "localhost"
    port = 8888
  }]
  property = [{
    name  = "name"
    value = "value"
  }]
}