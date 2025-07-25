
data "datapower_documentcryptomap" "test" {
  depends_on = [datapower_documentcryptomap.test]
  app_domain = "acc_test_domain"
}