
resource "datapower_crypto_certificate" "test" {
  id         = "ResTestCryptoCertificate"
  app_domain = "acceptance_test"
  filename   = "cert:///acc-test-server.crt"
}