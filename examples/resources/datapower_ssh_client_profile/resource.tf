
resource "datapower_ssh_client_profile" "test" {
  id            = "ResTestSSHClientProfile"
  app_domain    = "acceptance_test"
  user_name     = "someuser"
  profile_usage = "sftp"
}