
resource "datapower_oauth_supported_client_group" "test" {
  id         = "ResTestOAuthSupportedClientGroup"
  app_domain = "acceptance_test"
  client     = ["AccTest_OAuthSupportedClient"]
}