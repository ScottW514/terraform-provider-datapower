
resource "datapower_web_sphere_jms_server" "test" {
  id         = "ResTest_WebSphereJMSServer"
  app_domain = "acceptance_test"
  endpoint = [{
    host      = "localhost"
    port      = 8888
    transport = "TCP"
  }]
  messaging_bus = "bus"
}