
data "datapower_logtarget" "test" {
  depends_on = [datapower_logtarget.test]
  app_domain = "acc_test_domain"
}