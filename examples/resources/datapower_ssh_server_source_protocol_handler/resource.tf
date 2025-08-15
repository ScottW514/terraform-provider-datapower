
resource "datapower_ssh_server_source_protocol_handler" "test" {
  id                = "ResTestSSHServerSourceProtocolHandler"
  app_domain        = "acceptance_test"
  local_address     = "0.0.0.0"
  local_port        = 22
  default_directory = "/"
}