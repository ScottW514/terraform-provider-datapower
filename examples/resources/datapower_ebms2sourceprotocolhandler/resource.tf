
resource "datapower_ebms2sourceprotocolhandler" "test" {
  id            = "ResTestEBMS2SourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 80
}