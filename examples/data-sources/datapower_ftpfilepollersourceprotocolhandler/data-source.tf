
data "datapower_ftpfilepollersourceprotocolhandler" "test" {
  depends_on = [datapower_ftpfilepollersourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}