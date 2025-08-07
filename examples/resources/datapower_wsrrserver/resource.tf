
resource "datapower_wsrrserver" "test" {
  id         = "ResTest_WSRRServer"
  app_domain = "acceptance_test"
  soap_url   = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}