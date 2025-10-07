
resource "datapower_aaa_jwt_validator" "test" {
  id         = "ResTestAAAJWTValidator"
  app_domain = "acceptance_test"
  val_method = {
  }
  claims = [{
    value = "value"
  }]
  username_claim = "sub"
}