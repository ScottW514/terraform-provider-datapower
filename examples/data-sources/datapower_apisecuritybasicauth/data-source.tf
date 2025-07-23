
data "datapower_apisecuritybasicauth" "test" {
  depends_on = [datapower_apisecuritybasicauth.test]
  app_domain = "acc_test_domain"
}