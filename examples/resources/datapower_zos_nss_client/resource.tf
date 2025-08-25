
resource "datapower_zos_nss_client" "test" {
  id             = "ResTestZosNSSClient"
  app_domain     = "acceptance_test"
  remote_address = "remote.host"
  remote_port    = 443
  client_id      = "client_id"
  system_name    = "sysname"
  user_name      = "username"
  password_alias = "AccTest_PasswordAlias"
  ssl_client     = "AccTest_SSLClientProfile"
}