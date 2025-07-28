
resource "datapower_probe" "test" {
  app_domain      = "acc_test_domain"
  max_records     = 1000
  gateway_peering = "default-gateway-peering"
}