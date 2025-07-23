
data "datapower_oauthsupportedclient" "test" {
  depends_on = [datapower_oauthsupportedclient.test]
  app_domain = "acc_test_domain"
}