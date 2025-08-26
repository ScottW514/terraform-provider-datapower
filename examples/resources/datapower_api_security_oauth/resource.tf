
resource "datapower_api_security_oauth" "test" {
  id             = "ResTestAPISecurityOAuth"
  app_domain     = "acceptance_test"
  oauth_provider = "AccTest_OAuthProviderSettings"
  oauth_flow     = "implicit"
}