
resource "datapower_zosnssclient" "test" {
  id             = "ResTestZosNSSClient"
  app_domain     = "acceptance_test"
  remote_address = "remote.host"
  remote_port    = 443
  client_id      = "client_id"
  system_name    = "ResTestsystem"
  user_name      = "username"
}