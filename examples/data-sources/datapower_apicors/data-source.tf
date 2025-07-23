
data "datapower_apicors" "test" {
  depends_on = [datapower_apicors.test]
  app_domain = "acc_test_domain"
}