
resource "datapower_deployment_policy_parameters_binding" "test" {
  id         = "ResTestDeploymentPolicyParametersBinding"
  app_domain = "acceptance_test"
  deployment_policy_parameter = [{
    parameter_name  = "ResTestparameter"
    parameter_value = "parameter_value"
  }]
}