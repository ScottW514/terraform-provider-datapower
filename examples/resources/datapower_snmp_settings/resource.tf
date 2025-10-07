
resource "datapower_snmp_settings" "test" {
  local_port = 161
  policies = [{
    community = "private"
    domain    = "acc_test_domain"
    mode      = "none"
  }]
  policies_mq = null
  targets = [{
    host = "10.10.10.10"
  }]
  contexts = [{
    context = "context"
    domain  = "acc_test_domain"
  }]
  security_level = "authPriv"
  access_level   = "read-only"
}