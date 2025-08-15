
resource "datapower_luna_ha_group" "test" {
  id         = "ResTestLunaHAGroup"
  app_domain = "acceptance_test"
  group_name = "groupname"
  member     = ["AccTest_LunaPartition"]
}