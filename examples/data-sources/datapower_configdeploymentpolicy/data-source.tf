
data "datapower_configdeploymentpolicy" "test" {
  depends_on = [datapower_configdeploymentpolicy.test]
  app_domain = "acc_test_domain"
}