
resource "datapower_api_connect_gateway_service" "test" {
  app_domain      = "acceptance_test"
  local_address   = "0.0.0.0"
  local_port      = 3000
  gateway_peering = "default-gateway-peering"
  proxy_policy    = { proxy_policy_enable = false, remote_address = "localhost", remote_port = 8080 }
}