
data "datapower_policyattachments" "test" {
  depends_on = [datapower_policyattachments.test]
  app_domain = "acc_test_domain"
}