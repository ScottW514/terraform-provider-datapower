
resource "datapower_sshclientprofile" "test" {
  id            = "ResTestSSHClientProfile"
  app_domain    = "acceptance_test"
  user_name     = "someuser"
  profile_usage = "sftp"
}