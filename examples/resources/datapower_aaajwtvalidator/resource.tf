
resource "datapower_aaajwtvalidator" "test" {
  id             = "AAAJWTValidator_name"
  app_domain     = "acc_test_domain"
  username_claim = "sub"
}