
data "datapower_oauthprovidersettings" "test" {
  depends_on = [datapower_oauthprovidersettings.test]
  app_domain = "acc_test_domain"
}