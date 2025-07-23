
data "datapower_xmlfirewallservice" "test" {
  depends_on = [datapower_xmlfirewallservice.test]
  app_domain = "acc_test_domain"
}