
resource "datapower_webspherejmsserver" "test" {
  id         = "ResTest_WebSphereJMSServer"
  app_domain = "acceptance_test"
  endpoint = [{
    host      = "localhost"
    port      = 8888
    transport = "TCP"
  }]
  messaging_bus = "bus"
}