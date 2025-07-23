
data "datapower_apiclientidentification" "test" {
  depends_on = [datapower_apiclientidentification.test]
  app_domain = "acc_test_domain"
}