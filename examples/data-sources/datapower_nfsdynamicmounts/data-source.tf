
data "datapower_nfsdynamicmounts" "test" {
  depends_on = [datapower_nfsdynamicmounts.test]
  app_domain = "acc_test_domain"
}