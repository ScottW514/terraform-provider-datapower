
resource "datapower_usergroup" "test" {
  id              = "UserGroup_name"
  access_policies = ["*/*/*?Access=r"]
}