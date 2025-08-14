
resource "datapower_xslproxyservice" "test" {
  id             = "ResTestXSLProxyService"
  app_domain     = "acceptance_test"
  type           = "static-backend"
  xml_manager    = "default"
  local_port     = 8922
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}