
resource "datapower_load_balancer_group" "test" {
  id            = "ResTestLoadBalancerGroup"
  app_domain    = "acceptance_test"
  algorithm     = "round-robin"
  retrieve_info = false
  damp          = 120
  lb_group_members = [{
    server = "lbhost"
  }]
  lb_group_checks = {
    active    = false
    port      = 80
    ssl       = "off"
    timeout   = 10
    frequency = 180
  }
  lb_group_affinity_conf = {
    enable_sa = true
    insertion_attributes = {
    }
  }
}