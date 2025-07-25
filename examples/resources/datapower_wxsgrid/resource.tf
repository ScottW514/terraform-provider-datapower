
resource "datapower_wxsgrid" "test" {
  id             = "_name"
  app_domain     = "acc_test_domain"
  collective     = datapower_loadbalancergroup.test.id
  grid           = "gridname"
  user_name      = "username"
  password_alias = datapower_passwordalias.test.id
  timeout        = 1000
}