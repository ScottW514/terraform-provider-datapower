
data "datapower_wsrrsubscription" "test" {
  depends_on = [datapower_wsrrsubscription.test]
  app_domain = "acc_test_domain"
}