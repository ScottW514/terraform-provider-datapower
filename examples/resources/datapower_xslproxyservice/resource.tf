
resource "datapower_xslproxyservice" "test" {
  id             = "XSLProxyService_name"
  app_domain     = "acc_test_domain"
  type           = "static-backend"
  xml_manager    = "default"
  local_port     = 8899
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}