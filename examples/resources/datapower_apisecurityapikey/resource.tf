
resource "datapower_apisecurityapikey" "test" {
  id         = "APISecurityAPIKey_test"
  app_domain = "acc_test_domain"
  where      = "header"
  type       = "id"
  key_name   = "keyname"
}