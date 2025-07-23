
data "datapower_apisecuritytokenmanager" "test" {
  depends_on = [datapower_apisecuritytokenmanager.test]
  app_domain = "acc_test_domain"
}