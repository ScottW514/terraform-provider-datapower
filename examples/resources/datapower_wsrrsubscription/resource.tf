
resource "datapower_wsrrsubscription" "test" {
  id          = "ResTestWSRRSubscription"
  app_domain  = "acceptance_test"
  server      = "AccTest_WSRRServer"
  object_type = "wsdl"
  object_name = "ResTestobject"
}