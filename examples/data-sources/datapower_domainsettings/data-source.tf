
data "datapower_domainsettings" "test" {
  depends_on = [datapower_domainsettings.test]
  app_domain = "acc_test_domain"
}