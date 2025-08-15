
resource "datapower_snmp_settings" "test" {
  local_port     = 161
  security_level = "authPriv"
  access_level   = "read-only"
}