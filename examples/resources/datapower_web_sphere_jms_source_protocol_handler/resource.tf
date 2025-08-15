
resource "datapower_web_sphere_jms_source_protocol_handler" "test" {
  id         = "ResTestWebSphereJMSSourceProtocolHandler"
  app_domain = "acceptance_test"
  server     = "AccTest_WebSphereJMSServer"
  get_queue  = "getqueue"
}