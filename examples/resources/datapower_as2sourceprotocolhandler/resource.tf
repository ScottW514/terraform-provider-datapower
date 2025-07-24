
resource "datapower_as2sourceprotocolhandler" "test" {
  id            = "AS2SourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port    = 80
}