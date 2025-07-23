
data "datapower_cryptokey" "test" {
  depends_on = [datapower_cryptokey.test]
  app_domain = "acc_test_domain"
}