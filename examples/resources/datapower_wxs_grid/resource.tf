
resource "datapower_wxs_grid" "test" {
  id             = "ResTestWXSGrid"
  app_domain     = "acceptance_test"
  collective     = "AccTest_LoadBalancerGroup"
  grid           = "gridname"
  user_name      = "username"
  password_alias = "AccTest_PasswordAlias"
  timeout        = 1000
}