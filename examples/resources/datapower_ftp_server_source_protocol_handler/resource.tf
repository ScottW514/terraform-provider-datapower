
resource "datapower_ftp_server_source_protocol_handler" "test" {
  id            = "ResTestFTPServerSourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 21
}