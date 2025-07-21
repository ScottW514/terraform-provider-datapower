
resource "datapower_file" "test" {
  app_domain  = "acc_test_domain"
  remote_path = "cert:///test_file.txt"
  local_path  = "/workspaces/terraform-provider-datapower/test/test_file.txt"
}
