
data "datapower_cryptoidentcred" "test" {
  depends_on = [datapower_cryptoidentcred.test]
  app_domain = "acc_test_domain"
}