
data "datapower_aaajwtgenerator" "test" {
  depends_on = [datapower_aaajwtgenerator.test]
  app_domain = "acc_test_domain"
}