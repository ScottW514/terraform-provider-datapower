
resource "datapower_style_policy_action" "test" {
  id         = "ResTestStylePolicyAction"
  app_domain = "acceptance_test"
  type       = "xform"
  input      = "INPUT"
  transform  = "tfile"
  parse_settings_reference = {
    url = "some_url"
  }
  output        = "OUTPUT"
  named_inputs  = null
  named_outputs = null
  stylesheet_parameters = [{
    parameter_value = "PARAMETER-VALUE"
  }]
  condition = [{
    expression       = "*"
    condition_action = "*"
  }]
}