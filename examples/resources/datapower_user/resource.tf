
resource "datapower_user" "test" {
  id                = "User_name"
  password          = "Password$123"
  access_level      = "group-defined"
  group_name        = datapower_usergroup.test.id
  snmp_creds        = null
  hashed_snmp_creds = null
}