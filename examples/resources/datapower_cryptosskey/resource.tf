
resource "datapower_cryptosskey" "test" {
  id         = "ResTestCryptoSSKey"
  app_domain = "acceptance_test"
  filename   = "cert:///acc-test-server.key"
}