
resource "datapower_crypto_kerberos_kdc" "test" {
  id         = "ResTestCryptoKerberosKDC"
  app_domain = "acceptance_test"
  realm      = "realm"
  server     = "localhost"
}