
resource "datapower_apisecurityoauth" "test" {
  id          = "APISecurityOAuth_name"
  app_domain  = "acc_test_domain"
  o_auth_flow = "implicit"
}