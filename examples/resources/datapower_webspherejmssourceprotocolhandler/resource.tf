
resource "datapower_webspherejmssourceprotocolhandler" "test" {
  id         = "WebSphereJMSSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  server     = datapower_webspherejmsserver.test.id
  get_queue  = "getqueue"
}