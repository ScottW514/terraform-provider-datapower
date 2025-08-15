
resource "datapower_wsrr_subscription" "test" {
  id          = "ResTestWSRRSubscription"
  app_domain  = "acceptance_test"
  server      = "AccTest_WSRRServer"
  namespace   = "namespace"
  object_type = "wsdl"
  object_name = "ResTestobject"
}