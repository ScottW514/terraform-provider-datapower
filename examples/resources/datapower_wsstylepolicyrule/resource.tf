
resource "datapower_wsstylepolicyrule" "test" {
  id            = "WSStylePolicyRule_name"
  app_domain    = "acc_test_domain"
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}