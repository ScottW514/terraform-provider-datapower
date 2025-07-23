
data "datapower_sshdomainclientprofile" "test" {
  depends_on = [datapower_sshdomainclientprofile.test]
  app_domain = "acc_test_domain"
}