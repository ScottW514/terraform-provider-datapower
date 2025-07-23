
data "datapower_deploymentpolicyparametersbinding" "test" {
  depends_on = [datapower_deploymentpolicyparametersbinding.test]
  app_domain = "acc_test_domain"
}