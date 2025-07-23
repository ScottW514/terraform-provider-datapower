
data "datapower_sslsnimapping" "test" {
  depends_on = [datapower_sslsnimapping.test]
  app_domain = "acc_test_domain"
}