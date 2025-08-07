
resource "datapower_wsstylepolicyrule" "test" {
  id            = "ResTest_WSStylePolicyRule"
  app_domain    = "acceptance_test"
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}