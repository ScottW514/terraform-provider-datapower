
resource "datapower_oauth_provider_settings" "test" {
  id            = "ResTestOAuthProviderSettings"
  app_domain    = "acceptance_test"
  provider_type = "native"
  supported_grant_types = {
  }
  supported_client_types = {
  }
  api_c_token_secret = "AccTest_CryptoSSKey"
  adv_scope_url_security = {
  }
  api_c_oidc_hybrid_response_types = {
  }
  metadata_from = {
  }
  external_revocation_url_security = {
  }
  third_party_introspect_url_security = {
  }
}