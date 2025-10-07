
resource "datapower_config_deployment_policy" "test" {
  id         = "ResTestConfigDeploymentPolicy"
  app_domain = "acceptance_test"
  modified_config = [{
    match = "10.10.1.1/domainA/services/xslproxy?Value=myhost"
    type  = "change"
    value = "somevalue"
  }]
}