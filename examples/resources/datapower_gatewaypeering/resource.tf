
resource "datapower_gatewaypeering" "test" {
  id         = "ResTestGatewayPeering"
  app_domain = "acceptance_test"
  local_port = 16380
}