
resource "datapower_tam" "test" {
  id                 = "ResTestTAM"
  app_domain         = "acceptance_test"
  configuration_file = "local:///tam.config"
  ssl_key_file       = "cert:///ssl.key"
  ssl_key_stash_file = "cert:///ssl.stash"
  use_local_mode     = false
  tam_use_basic_user = false
}