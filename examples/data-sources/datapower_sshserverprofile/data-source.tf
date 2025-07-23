
data "datapower_sshserverprofile" "test" {
  depends_on = [datapower_sshserverprofile.test]
  app_domain = "acc_test_domain"
}