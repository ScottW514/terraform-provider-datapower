
resource "datapower_wxsgrid" "test" {
  id             = "WXSGrid_name"
  app_domain     = "acc_test_domain"
  collective     = "TestAccLoadBalancerGroup"
  grid           = "gridname"
  user_name      = "username"
  password_alias = "TestAccPasswordAlias"
  timeout        = 1000
}