
data "datapower_cryptosskey" "test" {
  depends_on = [datapower_cryptosskey.test]
  app_domain = "acc_test_domain"
}