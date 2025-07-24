
data "datapower_as1pollersourceprotocolhandler" "test" {
  depends_on = [datapower_as1pollersourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}