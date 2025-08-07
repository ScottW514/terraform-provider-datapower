
resource "datapower_sshserversourceprotocolhandler" "test" {
  id                = "ResTestSSHServerSourceProtocolHandler"
  app_domain        = "acceptance_test"
  local_address     = "0.0.0.0"
  local_port        = 22
  default_directory = "/"
}