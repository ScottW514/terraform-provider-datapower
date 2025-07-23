
data "datapower_apisecurityapikey" "test" {
  depends_on = [datapower_apisecurityapikey.test]
  app_domain = "acc_test_domain"
}