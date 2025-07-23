
data "datapower_apisecurityoauth" "test" {
  depends_on = [datapower_apisecurityoauth.test]
  app_domain = "acc_test_domain"
}