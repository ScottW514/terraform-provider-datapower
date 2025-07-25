
resource "datapower_cryptokerberoskdc" "test" {
  id         = "CryptoKerberosKDC_name"
  app_domain = "acc_test_domain"
  realm      = "realm"
  server     = "localhost"
}