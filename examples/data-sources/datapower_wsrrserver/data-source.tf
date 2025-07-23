
data "datapower_wsrrserver" "test" {
  depends_on = [datapower_wsrrserver.test]
  app_domain = "acc_test_domain"
}