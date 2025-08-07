
resource "datapower_file" "test" {
  app_domain  = "acc_test_domain"
  remote_path = "cert:///test_file.txt"
  local_path  = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
  dependency_actions = [
    {
      target_type   = "resource_datapower_domain",
      target_domain = "acc_test_domain",
      target_id     = "ignored",
      action        = "quiesce"
      on_create     = true
    }
  ]
}
