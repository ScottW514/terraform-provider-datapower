
data "datapower_xslcoprocservice" "test" {
  depends_on = [datapower_xslcoprocservice.test]
  app_domain = "acc_test_domain"
}