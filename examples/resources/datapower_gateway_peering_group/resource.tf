
resource "datapower_gateway_peering_group" "test" {
  id         = "ResTestGatewayPeeringGroup"
  app_domain = "acceptance_test"
  mode       = "peer"
  peer_nodes = [{
    host     = "10.10.10.10"
    priority = 100
  }]
  cluster_nodes = [{
    host = "localhost"
  }]
  enable_ssl = false
}