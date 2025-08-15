
resource "datapower_gateway_peering_manager" "test" {
  app_domain                  = "acceptance_test"
  api_connect_gateway_service = "default-gateway-peering"
  rate_limit                  = "default-gateway-peering"
  subscription                = "default-gateway-peering"
  ratelimit_module            = "default-gateway-peering"
}