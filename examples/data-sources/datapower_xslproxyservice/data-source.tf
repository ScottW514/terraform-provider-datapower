
data "datapower_xslproxyservice" "test" {
  depends_on = [datapower_xslproxyservice.test]
  app_domain = "acc_test_domain"
}