
resource "datapower_zosnssclient" "test" {
  id             = "ZosNSSClient_name"
  app_domain     = "acc_test_domain"
  remote_address = "remote.host"
  remote_port    = 443
  client_id      = "client_id"
  system_name    = "system_name"
  user_name      = "username"
}