
data "datapower_cryptokerberoskdc" "test" {
  depends_on = [datapower_cryptokerberoskdc.test]
  app_domain = "acc_test_domain"
}