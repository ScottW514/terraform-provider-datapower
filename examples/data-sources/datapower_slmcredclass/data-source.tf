
data "datapower_slmcredclass" "test" {
  depends_on = [datapower_slmcredclass.test]
  app_domain = "acc_test_domain"
}