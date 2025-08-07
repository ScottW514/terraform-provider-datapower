
resource "datapower_webspherejmssourceprotocolhandler" "test" {
  id         = "ResTestWebSphereJMSSourceProtocolHandler"
  app_domain = "acceptance_test"
  server     = "AccTest_WebSphereJMSServer"
  get_queue  = "getqueue"
}