
resource "datapower_cryptoidentcred" "test" {
  id          = "CookieAttributePolicy_name"
  app_domain  = "acc_test_domain"
  key         = "TestAccCryptoKey"
  certificate = "TestAccCryptoCertificate"
}