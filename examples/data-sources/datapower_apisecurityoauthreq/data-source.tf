
data "datapower_apisecurityoauthreq" "test" {
  depends_on = [datapower_apisecurityoauthreq.test]
  app_domain = "acc_test_domain"
}