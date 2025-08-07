
resource "datapower_loadbalancergroup" "test" {
  id            = "ResTestLoadBalancerGroup"
  app_domain    = "acceptance_test"
  algorithm     = "round-robin"
  retrieve_info = false
  damp          = 120
}