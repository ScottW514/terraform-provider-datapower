
resource "datapower_oauth_provider_settings" "test" {
  id                = "ResTestOAuthProviderSettings"
  app_domain        = "acceptance_test"
  provider_type     = "native"
  apic_token_secret = "AccTest_CryptoSSKey"
}