
resource "datapower_sshclientprofile" "test" {
  id            = "SSHClientProfile_name"
  app_domain    = "acc_test_domain"
  user_name     = "someuser"
  profile_usage = "sftp"
}