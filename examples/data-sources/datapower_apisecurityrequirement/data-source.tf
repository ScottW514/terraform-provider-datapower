
data "datapower_apisecurityrequirement" "test" {
  depends_on = [datapower_apisecurityrequirement.test]
  app_domain = "acc_test_domain"
}