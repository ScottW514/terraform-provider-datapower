
resource "datapower_webspherejmsserver" "test" {
  id         = "WebSphereJMSServer_name"
  app_domain = "acc_test_domain"
  endpoint = [{
    host      = "localhost"
    port      = 8888
    transport = "TCP"
  }]
  messaging_bus = "bus"
}