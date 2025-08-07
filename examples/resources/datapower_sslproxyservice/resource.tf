
resource "datapower_sslproxyservice" "test" {
  id             = "ResTestSSLProxyService"
  app_domain     = "acceptance_test"
  local_port     = 8888
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}