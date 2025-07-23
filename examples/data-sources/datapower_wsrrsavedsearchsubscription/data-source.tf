
data "datapower_wsrrsavedsearchsubscription" "test" {
  depends_on = [datapower_wsrrsavedsearchsubscription.test]
  app_domain = "acc_test_domain"
}