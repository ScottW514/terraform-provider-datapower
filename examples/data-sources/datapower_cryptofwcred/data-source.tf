
data "datapower_cryptofwcred" "test" {
  depends_on = [datapower_cryptofwcred.test]
  app_domain = "acc_test_domain"
}