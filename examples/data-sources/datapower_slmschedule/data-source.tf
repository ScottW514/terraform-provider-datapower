
data "datapower_slmschedule" "test" {
  depends_on = [datapower_slmschedule.test]
  app_domain = "acc_test_domain"
}