
resource "datapower_mqmanager" "test" {
  id            = "ResTestMQManager"
  app_domain    = "acceptance_test"
  host_name     = "localhost"
  cache_timeout = 60
  xml_manager   = "default"
}