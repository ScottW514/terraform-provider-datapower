
resource "datapower_api_security_api_key" "test" {
  id         = "ResTestAPISecurityAPIKey"
  app_domain = "acceptance_test"
  where      = "header"
  type       = "id"
  key_name   = "keyname"
}