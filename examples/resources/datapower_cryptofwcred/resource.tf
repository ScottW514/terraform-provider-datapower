
resource "datapower_cryptofwcred" "test" {
  id          = "ResTestCryptoFWCred"
  app_domain  = "acceptance_test"
  private_key = ["AccTest_CryptoKey"]
}