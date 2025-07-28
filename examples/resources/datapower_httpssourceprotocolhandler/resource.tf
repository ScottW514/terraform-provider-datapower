
resource "datapower_httpssourceprotocolhandler" "test" {
  id                     = "HTTPSSourceProtocolHandler_name"
  app_domain             = "acc_test_domain"
  ssl_server_config_type = "server"
}