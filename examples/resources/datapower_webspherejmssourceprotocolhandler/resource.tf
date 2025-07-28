
resource "datapower_webspherejmssourceprotocolhandler" "test" {
  id         = "WebSphereJMSSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  server     = "TestAccWebSphereJMSServer"
  get_queue  = "getqueue"
}