
resource "datapower_mqmanager" "test" {
  id            = "MQManager_name"
  app_domain    = "acc_test_domain"
  host_name     = "localhost"
  cache_timeout = 60
  xml_manager   = "default"
}