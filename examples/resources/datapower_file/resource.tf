
resource "datapower_file" "acc_test" {
  app_domain  = "acceptance_test"
  remote_path = "cert:///test_file.txt"
  local_path  = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
}
