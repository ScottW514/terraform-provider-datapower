
data "datapower_webtokenservice" "test" {
  depends_on = [datapower_webtokenservice.test]
  app_domain = "acc_test_domain"
}