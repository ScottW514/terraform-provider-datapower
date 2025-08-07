
resource "datapower_user" "test" {
  id                = "ResTest_User"
  password          = "Password$123"
  access_level      = "group-defined"
  group_name        = "AccTest_UserGroup"
  snmp_creds        = null
  hashed_snmp_creds = null
}