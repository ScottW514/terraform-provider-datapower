
resource "datapower_crypto_val_cred" "test" {
  id          = "ResTestCryptoValCred"
  app_domain  = "acceptance_test"
  certificate = ["AccTest_CryptoCertificate"]
}