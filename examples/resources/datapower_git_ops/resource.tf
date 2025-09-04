
resource "datapower_git_ops" "test" {
  app_domain        = "acceptance_test"
  connection_type   = "https"
  mode              = "read-write"
  commit_identifier = "main"
  remote_location   = "https://github.com/ScottW514/terraform-provider-datapower"
  username          = "gitusername"
  password          = "AccTest_PasswordAlias"
  tls_valcred       = "AccTest_CryptoValCred"
  git_user          = "Git User"
  git_email         = "git@user.domain"
}