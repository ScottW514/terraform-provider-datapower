
resource "datapower_sshserversourceprotocolhandler" "test" {
  id                = "SSHServerSourceProtocolHandler_name"
  app_domain        = "acc_test_domain"
  local_address     = "0.0.0.0"
  local_port        = 22
  default_directory = "/"
}