
data "datapower_webspherejmssourceprotocolhandler" "test" {
  depends_on = [datapower_webspherejmssourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}