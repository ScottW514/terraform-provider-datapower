
resource "datapower_as3_source_protocol_handler" "test" {
  id            = "ResTestAS3SourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 21
}