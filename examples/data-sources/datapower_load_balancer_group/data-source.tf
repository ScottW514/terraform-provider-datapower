
data "datapower_load_balancer_group" "test" {
  id         = "AccTest_LoadBalancerGroup"
  app_domain = "acceptance_test"
}