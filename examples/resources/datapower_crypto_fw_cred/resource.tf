
resource "datapower_crypto_fw_cred" "test" {
  id          = "ResTestCryptoFWCred"
  app_domain  = "acceptance_test"
  private_key = ["AccTest_CryptoKey"]
}