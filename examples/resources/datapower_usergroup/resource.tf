
resource "datapower_usergroup" "test" {
  id              = "_UserGroup_name"
  access_policies = ["*/*/*?Access=r"]
}