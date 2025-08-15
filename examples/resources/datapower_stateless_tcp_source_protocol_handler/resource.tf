
resource "datapower_stateless_tcp_source_protocol_handler" "test" {
  id            = "ResTestStatelessTCPSourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 4000
}