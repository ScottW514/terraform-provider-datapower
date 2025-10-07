
resource "datapower_policy_parameters" "test" {
  id         = "ResTestPolicyParameters"
  app_domain = "acceptance_test"
  policy_parameter = [{
    parameter_name  = "PolicyParameterName"
    parameter_value = "PolicyParameterValue"
  }]
}