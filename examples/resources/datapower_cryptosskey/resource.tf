
resource "datapower_cryptosskey" "test" {
  id         = "CryptoSSKey_name"
  app_domain = "acc_test_domain"
  filename   = "cert:///acc-test-server.key"
}