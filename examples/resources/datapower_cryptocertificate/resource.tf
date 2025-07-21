
resource "datapower_cryptocertificate" "test" {
  id         = "CryptoCertificate_test"
  app_domain = "acc_test_domain"
  filename   = "cert:///acc-test-server.crt"
}