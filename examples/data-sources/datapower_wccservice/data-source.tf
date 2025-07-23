
data "datapower_wccservice" "test" {
  depends_on = [datapower_wccservice.test]
  app_domain = "acc_test_domain"
}