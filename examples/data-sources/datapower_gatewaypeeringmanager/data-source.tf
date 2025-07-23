
data "datapower_gatewaypeeringmanager" "test" {
  depends_on = [datapower_gatewaypeeringmanager.test]
  app_domain = "acc_test_domain"
}