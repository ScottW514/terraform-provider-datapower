
resource "datapower_stylepolicyrule" "test" {
  id            = "ResTestStylePolicyRule"
  app_domain    = "acceptance_test"
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}