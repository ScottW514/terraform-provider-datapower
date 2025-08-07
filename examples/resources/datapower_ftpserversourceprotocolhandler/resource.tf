
resource "datapower_ftpserversourceprotocolhandler" "test" {
  id            = "ResTestFTPServerSourceProtocolHandler"
  app_domain    = "acceptance_test"
  local_address = "0.0.0.0"
  local_port    = 21
}