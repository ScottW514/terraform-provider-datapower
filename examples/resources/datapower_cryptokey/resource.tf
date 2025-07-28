
resource "datapower_cryptokey" "test" {
  id         = "CryptoKey_name"
  app_domain = "acc_test_domain"
  filename   = "cert:///acc-test-server.key"
  alias      = "TestAccPasswordAlias"
}