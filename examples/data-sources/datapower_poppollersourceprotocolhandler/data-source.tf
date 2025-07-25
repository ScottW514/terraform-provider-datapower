
data "datapower_poppollersourceprotocolhandler" "test" {
  depends_on = [datapower_poppollersourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}