
data "datapower_visibilitylist" "test" {
  depends_on = [datapower_visibilitylist.test]
  app_domain = "acc_test_domain"
}