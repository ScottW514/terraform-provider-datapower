
resource "datapower_webservicemonitor" "test" {
  id           = "ResTestWebServiceMonitor"
  app_domain   = "acceptance_test"
  wsdlurl      = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}