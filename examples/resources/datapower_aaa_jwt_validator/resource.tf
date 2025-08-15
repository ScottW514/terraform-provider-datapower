
resource "datapower_aaa_jwt_validator" "test" {
  id             = "ResTestAAAJWTValidator"
  app_domain     = "acceptance_test"
  username_claim = "sub"
}