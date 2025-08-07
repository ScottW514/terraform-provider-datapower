
resource "datapower_as2proxysourceprotocolhandler" "test" {
  id                        = "ResTestAS2ProxySourceProtocolHandler"
  app_domain                = "acceptance_test"
  local_address             = "0.0.0.0"
  local_port                = 80
  remote_address            = "10.10.10.10"
  remote_port               = 8888
  remote_connection_timeout = 60
  xml_manager               = "default"
}