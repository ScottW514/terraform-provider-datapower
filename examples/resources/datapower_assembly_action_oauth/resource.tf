
resource "datapower_assembly_action_oauth" "test" {
  id                                = "ResTestAssemblyActionOAuth"
  app_domain                        = "acceptance_test"
  oauth_provider_settings_reference = { "Default" : "AccTest_OAuthProviderSettings" }
}