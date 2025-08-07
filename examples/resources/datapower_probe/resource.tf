
resource "datapower_probe" "test" {
  app_domain      = "acceptance_test"
  max_records     = 1000
  gateway_peering = "default-gateway-peering"
}