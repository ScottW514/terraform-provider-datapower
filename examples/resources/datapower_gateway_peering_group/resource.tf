
resource "datapower_gateway_peering_group" "test" {
  id         = "ResTestGatewayPeeringGroup"
  app_domain = "acceptance_test"
  mode       = "peer"
  enable_ssl = false
}