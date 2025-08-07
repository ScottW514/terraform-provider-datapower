
resource "datapower_lunahagroup" "test" {
  id         = "ResTestLunaHAGroup"
  app_domain = "acceptance_test"
  group_name = "groupname"
  member     = ["AccTest_LunaPartition"]
}