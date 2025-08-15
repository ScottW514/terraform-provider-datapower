
resource "datapower_distributed_variable" "test" {
  app_domain      = "acceptance_test"
  gateway_peering = "default-gateway-peering"
}