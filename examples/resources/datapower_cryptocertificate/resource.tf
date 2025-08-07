
resource "datapower_cryptocertificate" "test" {
  id         = "ResTestCryptoCertificate"
  app_domain = "acceptance_test"
  filename   = "cert:///acc-test-server.crt"
}