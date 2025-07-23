
data "datapower_sshclientprofile" "test" {
  depends_on = [datapower_sshclientprofile.test]
  app_domain = "acc_test_domain"
}