
resource "datapower_usergroup" "test" {
  id              = "ResTest_UserGroup"
  access_policies = ["*/*/*?Access=r"]
}