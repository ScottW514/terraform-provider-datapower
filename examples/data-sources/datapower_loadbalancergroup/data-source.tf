
data "datapower_loadbalancergroup" "test" {
  depends_on = [datapower_loadbalancergroup.test]
  app_domain = "acc_test_domain"
}