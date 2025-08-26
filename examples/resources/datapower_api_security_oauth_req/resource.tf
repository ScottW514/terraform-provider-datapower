
resource "datapower_api_security_oauth_req" "test" {
  id                     = "ResTestAPISecurityOAuthReq"
  app_domain             = "acceptance_test"
  api_security_oauth_def = "AccTest_APISecurityOAuth"
}