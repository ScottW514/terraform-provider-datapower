
data "datapower_samlattributes" "test" {
  depends_on = [datapower_samlattributes.test]
  app_domain = "acc_test_domain"
}