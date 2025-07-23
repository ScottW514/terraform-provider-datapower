
resource "datapower_gitops" "test" {
  app_domain             = "acc_test_domain"
  connection_type        = "https"
  mode                   = "read-write"
  commit_identifier_type = "branch"
  commit_identifier      = "main"
  remote_location        = "https://github.com/ScottW514/terraform-provider-datapower"
}