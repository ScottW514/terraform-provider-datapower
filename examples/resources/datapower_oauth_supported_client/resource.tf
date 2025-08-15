
resource "datapower_oauth_supported_client" "test" {
  id                     = "ResTestOAuthSupportedClient"
  app_domain             = "acceptance_test"
  o_auth_role            = { "azsvr" : true }
  az_grant               = { "code" : true }
  generate_client_secret = false
  client_secret          = "secret"
  token_secret           = "AccTest_CryptoSSKey"
}