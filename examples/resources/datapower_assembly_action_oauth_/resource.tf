
resource "datapower_assembly_action_oauth_" "test" {
  id                                 = "ResTestAssemblyActionOAuth"
  app_domain                         = "acceptance_test"
  o_auth_provider_settings_reference = { "Default" : "AccTest_OAuthProviderSettings" }
}