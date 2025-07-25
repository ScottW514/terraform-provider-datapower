
data "datapower_statelesstcpsourceprotocolhandler" "test" {
  depends_on = [datapower_statelesstcpsourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}