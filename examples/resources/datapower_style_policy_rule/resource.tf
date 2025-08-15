
resource "datapower_style_policy_rule" "test" {
  id            = "ResTestStylePolicyRule"
  app_domain    = "acceptance_test"
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}