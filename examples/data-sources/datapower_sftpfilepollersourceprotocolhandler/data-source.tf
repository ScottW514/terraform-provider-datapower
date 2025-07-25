
data "datapower_sftpfilepollersourceprotocolhandler" "test" {
  depends_on = [datapower_sftpfilepollersourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}