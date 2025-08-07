
resource "datapower_cryptokerberoskdc" "test" {
  id         = "ResTestCryptoKerberosKDC"
  app_domain = "acceptance_test"
  realm      = "realm"
  server     = "localhost"
}