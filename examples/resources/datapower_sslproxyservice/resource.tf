
resource "datapower_sslproxyservice" "test" {
  id             = "ResTestSSLProxyService"
  app_domain     = "acceptance_test"
  local_port     = 4521
  remote_address = "10.10.10.10"
  remote_port    = 9999
  ssl_server     = "AccTest_SSLServerProfile"
  local_address  = "0.0.0.0"
}