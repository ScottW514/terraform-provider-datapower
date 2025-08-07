
resource "datapower_cryptoidentcred" "test" {
  id          = "ResTestCryptoIdentCred"
  app_domain  = "acceptance_test"
  key         = "AccTest_CryptoKey"
  certificate = "AccTest_CryptoCertificate"
}