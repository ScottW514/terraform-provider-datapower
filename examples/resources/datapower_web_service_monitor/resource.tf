
resource "datapower_web_service_monitor" "test" {
  id           = "ResTestWebServiceMonitor"
  app_domain   = "acceptance_test"
  wsdl_url     = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}