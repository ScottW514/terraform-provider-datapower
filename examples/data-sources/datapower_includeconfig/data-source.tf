
data "datapower_includeconfig" "test" {
  depends_on = [datapower_includeconfig.test]
  app_domain = "acc_test_domain"
}