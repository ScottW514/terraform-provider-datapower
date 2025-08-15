
resource "datapower_user_group" "test" {
  id              = "ResTest_UserGroup"
  access_policies = ["*/*/*?Access=r"]
}