
resource "datapower_aaajwtvalidator" "test" {
  id             = "AAAJWTValidator_test"
  app_domain     = "acc_test_domain"
  username_claim = "sub"
}