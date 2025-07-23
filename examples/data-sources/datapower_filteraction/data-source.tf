
data "datapower_filteraction" "test" {
  depends_on = [datapower_filteraction.test]
  app_domain = "acc_test_domain"
}