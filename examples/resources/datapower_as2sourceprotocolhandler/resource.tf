
resource "datapower_as2sourceprotocolhandler" "test" {
  id            = "ResTestAS2SourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 80
}