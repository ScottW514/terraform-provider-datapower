
resource "datapower_gatewaypeeringmanager" "test" {
  app_domain                  = "acc_test_domain"
  api_connect_gateway_service = "default-gateway-peering"
  rate_limit                  = "default-gateway-peering"
  subscription                = "default-gateway-peering"
  ratelimit_module            = "default-gateway-peering"
}