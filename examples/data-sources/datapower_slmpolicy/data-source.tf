
data "datapower_slmpolicy" "test" {
  depends_on = [datapower_slmpolicy.test]
  app_domain = "acc_test_domain"
}