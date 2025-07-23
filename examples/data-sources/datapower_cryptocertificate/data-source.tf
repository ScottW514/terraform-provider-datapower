
data "datapower_cryptocertificate" "test" {
  depends_on = [datapower_cryptocertificate.test]
  app_domain = "acc_test_domain"
}