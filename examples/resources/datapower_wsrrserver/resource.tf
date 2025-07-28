
resource "datapower_wsrrserver" "test" {
  id         = "_WSRRServer_name"
  app_domain = "acc_test_domain"
  soap_url   = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}