
resource "datapower_stylepolicyrule" "test" {
  id            = "StylePolicyRule_name"
  app_domain    = "acc_test_domain"
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}