
resource "datapower_ftpserversourceprotocolhandler" "test" {
  id            = "FTPServerSourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port    = 21
}