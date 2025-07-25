
resource "datapower_ebms3sourceprotocolhandler" "test" {
  id            = "EBMS3SourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port    = 80
}