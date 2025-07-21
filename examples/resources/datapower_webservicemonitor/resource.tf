
resource "datapower_webservicemonitor" "test" {
  id           = "WebServiceMonitor_name"
  app_domain   = "acc_test_domain"
  wsdlurl      = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}