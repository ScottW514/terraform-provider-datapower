
resource "datapower_user" "test" {
  id                  = "ResTest_User"
  password_wo         = "Password$123"
  password_wo_version = 1
  access_level        = "group-defined"
  group_name          = "AccTest_UserGroup"
  snmp_creds          = null
  hashed_snmp_creds   = null
}