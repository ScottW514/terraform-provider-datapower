
data "datapower_cryptovalcred" "test" {
  depends_on = [datapower_cryptovalcred.test]
  app_domain = "acc_test_domain"
}