
data "datapower_slmaction" "test" {
  depends_on = [datapower_slmaction.test]
  app_domain = "acc_test_domain"
}