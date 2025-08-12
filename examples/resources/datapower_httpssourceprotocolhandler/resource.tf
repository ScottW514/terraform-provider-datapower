
resource "datapower_httpssourceprotocolhandler" "test" {
  id                     = "ResTestHTTPSSourceProtocolHandler"
  app_domain             = "acceptance_test"
  ssl_server_config_type = "server"
  ssl_server             = "AccTest_SSLServerProfile"
}