
resource "datapower_cryptokey" "test" {
  id         = "CryptoKey_test"
  app_domain = "acc_test_domain"
  filename   = "cert:///acc-test-server.key"
  alias      = datapower_passwordalias.test.id
}