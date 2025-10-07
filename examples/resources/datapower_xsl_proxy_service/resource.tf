
resource "datapower_xsl_proxy_service" "test" {
  id             = "ResTestXSLProxyService"
  app_domain     = "acceptance_test"
  type           = "static-backend"
  xml_manager    = "default"
  local_port     = 8922
  remote_address = "10.10.10.10"
  remote_port    = 9999
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