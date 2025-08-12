
resource "datapower_gatewaypeering" "test" {
  id            = "ResTestGatewayPeering"
  app_domain    = "acceptance_test"
  local_address = "5.5.5.5"
  local_port    = 16380
  enable_ssl    = false
}