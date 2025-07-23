
data "datapower_aaajwtvalidator" "test" {
  depends_on = [datapower_aaajwtvalidator.test]
  app_domain = "acc_test_domain"
}