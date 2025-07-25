
data "datapower_wsendpointrewritepolicy" "test" {
  depends_on = [datapower_wsendpointrewritepolicy.test]
  app_domain = "acc_test_domain"
}