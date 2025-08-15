
resource "datapower_git_ops" "test" {
  app_domain             = "acceptance_test"
  connection_type        = "https"
  mode                   = "read-write"
  commit_identifier_type = "branch"
  commit_identifier      = "main"
  remote_location        = "https://github.com/ScottW514/terraform-provider-datapower"
}