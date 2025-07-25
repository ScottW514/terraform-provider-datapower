
resource "datapower_gatewaypeeringgroup" "test" {
  id         = "GatewayPeeringGroup_name"
  app_domain = "acc_test_domain"
  mode       = "peer"
  enable_ssl = true
}