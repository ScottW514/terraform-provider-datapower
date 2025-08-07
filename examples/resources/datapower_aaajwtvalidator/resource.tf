
resource "datapower_aaajwtvalidator" "test" {
  id             = "ResTestAAAJWTValidator"
  app_domain     = "acceptance_test"
  username_claim = "sub"
}