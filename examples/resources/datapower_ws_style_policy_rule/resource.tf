
resource "datapower_ws_style_policy_rule" "test" {
  id            = "ResTest_WSStylePolicyRule"
  app_domain    = "acceptance_test"
  actions       = ["AccTest_StylePolicyAction"]
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}