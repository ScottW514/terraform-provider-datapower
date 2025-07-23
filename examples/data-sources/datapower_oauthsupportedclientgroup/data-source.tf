
data "datapower_oauthsupportedclientgroup" "test" {
  depends_on = [datapower_oauthsupportedclientgroup.test]
  app_domain = "acc_test_domain"
}