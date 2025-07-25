
resource "datapower_lunahagroup" "test" {
  id         = "LunaHAGroup_name"
  app_domain = "acc_test_domain"
  group_name = "groupname"
  member     = [datapower_lunapartition.test.id]
}