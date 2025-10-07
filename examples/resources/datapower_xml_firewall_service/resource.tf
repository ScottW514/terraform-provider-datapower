
resource "datapower_xml_firewall_service" "test" {
  id          = "ResTestXMLFirewallService"
  app_domain  = "acceptance_test"
  type        = "dynamic-backend"
  xml_manager = "default"
  local_port  = 7575
  http_version = {
  }
  header_injection = [{
    header_tag_value = "SomeHeaderValue"
  }]
  header_suppression = null
  stylesheet_parameters = [{
    parameter_value = "PARAMETER-VALUE"
  }]
  debug_trigger = null
  local_address = "0.0.0.0"
}