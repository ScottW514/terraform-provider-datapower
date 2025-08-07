
resource "datapower_apisecurityoauthreq" "test" {
  id                      = "ResTestAPISecurityOAuthReq"
  app_domain              = "acceptance_test"
  api_security_o_auth_def = "AccTest_APISecurityOAuth"
}