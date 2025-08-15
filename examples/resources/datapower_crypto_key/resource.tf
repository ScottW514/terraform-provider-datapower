
resource "datapower_crypto_key" "test" {
  id         = "ResTestCryptoKey"
  app_domain = "acceptance_test"
  filename   = "cert:///acc-test-server.key"
  alias      = "AccTest_PasswordAlias"
}