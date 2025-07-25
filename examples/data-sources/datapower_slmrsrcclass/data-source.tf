
data "datapower_slmrsrcclass" "test" {
  depends_on = [datapower_slmrsrcclass.test]
  app_domain = "acc_test_domain"
}