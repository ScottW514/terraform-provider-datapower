
resource "datapower_crypto_sskey" "test" {
  id         = "ResTestCryptoSSKey"
  app_domain = "acceptance_test"
  filename   = "cert://tokensecret"
}