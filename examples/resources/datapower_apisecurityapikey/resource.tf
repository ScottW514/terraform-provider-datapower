
resource "datapower_apisecurityapikey" "test" {
  id         = "ResTestAPISecurityAPIKey"
  app_domain = "acceptance_test"
  where      = "header"
  type       = "id"
  key_name   = "keyname"
}