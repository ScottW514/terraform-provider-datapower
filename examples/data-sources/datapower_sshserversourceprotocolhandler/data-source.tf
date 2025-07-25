
data "datapower_sshserversourceprotocolhandler" "test" {
  depends_on = [datapower_sshserversourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}