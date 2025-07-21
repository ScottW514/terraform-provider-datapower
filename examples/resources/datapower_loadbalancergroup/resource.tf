
resource "datapower_loadbalancergroup" "test" {
  id            = "LoadBalancerGroup_name"
  app_domain    = "acc_test_domain"
  algorithm     = "round-robin"
  retrieve_info = false
  damp          = 120
}