
resource "datapower_cryptoidentcred" "test" {
  id          = "CookieAttributePolicy_name"
  app_domain  = "acc_test_domain"
  key         = datapower_cryptokey.test.id
  certificate = datapower_cryptocertificate.test.id
}