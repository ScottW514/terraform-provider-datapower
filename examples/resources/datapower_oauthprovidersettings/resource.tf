
resource "datapower_oauthprovidersettings" "test" {
  id            = "OAuthProviderSettings_name"
  app_domain    = "acc_test_domain"
  provider_type = "native"
}