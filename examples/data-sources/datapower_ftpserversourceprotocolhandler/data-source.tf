
data "datapower_ftpserversourceprotocolhandler" "test" {
  depends_on = [datapower_ftpserversourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}