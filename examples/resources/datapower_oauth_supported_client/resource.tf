
resource "datapower_oauth_supported_client" "test" {
  id                       = "ResTestOAuthSupportedClient"
  app_domain               = "acceptance_test"
  oauth_role               = { "azsvr" : true }
  az_grant                 = { "code" : true }
  generate_client_secret   = false
  client_secret_wo         = "secret"
  client_secret_wo_version = 1
  redirect_uri             = ["^https://example.com/redirect$"]
  scope                    = "*"
  token_secret             = "AccTest_CryptoSSKey"
}