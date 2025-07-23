
data "datapower_policyparameters" "test" {
  depends_on = [datapower_policyparameters.test]
  app_domain = "acc_test_domain"
}