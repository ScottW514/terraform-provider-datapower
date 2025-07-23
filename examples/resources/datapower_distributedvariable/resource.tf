
resource "datapower_distributedvariable" "test" {
  app_domain      = "acc_test_domain"
  gateway_peering = datapower_gatewaypeering.test.id
}