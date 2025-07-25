
data "datapower_nfsfilepollersourceprotocolhandler" "test" {
  depends_on = [datapower_nfsfilepollersourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}