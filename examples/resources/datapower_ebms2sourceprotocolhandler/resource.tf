
resource "datapower_ebms2sourceprotocolhandler" "test" {
  id            = "EBMS2SourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port    = 80
}