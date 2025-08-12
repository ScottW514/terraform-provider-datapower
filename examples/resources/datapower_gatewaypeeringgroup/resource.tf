
resource "datapower_gatewaypeeringgroup" "test" {
  id         = "ResTestGatewayPeeringGroup"
  app_domain = "acceptance_test"
  mode       = "peer"
  enable_ssl = false
}