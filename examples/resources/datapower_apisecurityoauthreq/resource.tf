
resource "datapower_apisecurityoauthreq" "test" {
  id                      = "APISecurityOAuthReq_test"
  app_domain              = "acc_test_domain"
  api_security_o_auth_def = datapower_apisecurityoauth.test.id
}