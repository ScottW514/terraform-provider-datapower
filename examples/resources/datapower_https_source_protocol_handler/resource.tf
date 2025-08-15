
resource "datapower_https_source_protocol_handler" "test" {
  id                     = "ResTestHTTPSSourceProtocolHandler"
  app_domain             = "acceptance_test"
  ssl_server_config_type = "server"
  ssl_server             = "AccTest_SSLServerProfile"
}