
resource "datapower_tam" "test" {
  id                 = "ResTestTAM"
  app_domain         = "acceptance_test"
  configuration_file = "local:///tam.config"
  ssl_key_file       = "cert:///ssl.key"
  ssl_key_stash_file = "cert:///ssl.stash"
  use_local_mode     = false
  tam_use_basic_user = false
  tam_fed_dirs = [{
    fed_name = "ResTestfed"
    suffix   = "suffix"
    host     = "ldap.host"
    port     = 389
    bind_dn  = "dn"
    bind_pw  = "AccTest_PasswordAlias"
    use_ssl  = false
  }]
  tam_az_replicas = [{
    tam_az_replica        = "replica"
    tam_az_replica_port   = 7136
    tam_az_replica_weight = 10
  }]
  tam_ras_trace = {
    tam_trace_enable    = false
    ldap_trace_enable   = false
    gs_kit_trace_enable = false
  }
}