
resource "datapower_api_security_oauth" "test" {
  id              = "ResTestAPISecurityOAuth"
  app_domain      = "acceptance_test"
  o_auth_provider = "AccTest_OAuthProviderSettings"
  o_auth_flow     = "implicit"
}