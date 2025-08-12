
resource "datapower_assemblyactionoauth" "test" {
  id                                 = "ResTestAssemblyActionOAuth"
  app_domain                         = "acceptance_test"
  o_auth_provider_settings_reference = { "Default" : "AccTest_OAuthProviderSettings" }
}