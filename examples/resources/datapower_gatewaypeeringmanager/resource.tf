
resource "datapower_gatewaypeeringmanager" "test" {
  app_domain                  = "acc_test_domain"
  api_connect_gateway_service = datapower_gatewaypeering.test.id
  rate_limit                  = datapower_gatewaypeering.test.id
  subscription                = datapower_gatewaypeering.test.id
  ratelimit_module            = datapower_gatewaypeering.test.id
}