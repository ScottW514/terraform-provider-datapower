
resource "datapower_wsrrsubscription" "test" {
  id          = "WSRRSubscription_name"
  app_domain  = "acc_test_domain"
  server      = "TestAccWSRRServer"
  object_type = "wsdl"
  object_name = "object_name"
}