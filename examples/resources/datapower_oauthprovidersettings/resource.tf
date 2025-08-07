
resource "datapower_oauthprovidersettings" "test" {
  id            = "ResTestOAuthProviderSettings"
  app_domain    = "acceptance_test"
  provider_type = "native"
}