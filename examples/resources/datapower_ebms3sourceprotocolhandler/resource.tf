
resource "datapower_ebms3sourceprotocolhandler" "test" {
  id            = "ResTestEBMS3SourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 80
}