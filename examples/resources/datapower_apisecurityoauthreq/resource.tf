
resource "datapower_apisecurityoauthreq" "test" {
  id                      = "APISecurityOAuthReq_name"
  app_domain              = "acc_test_domain"
  api_security_o_auth_def = "TestAccAPISecurityOAuth"
}