
resource "datapower_http_source_protocol_handler" "test" {
  id            = "ResTestHTTPSourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 8088
}