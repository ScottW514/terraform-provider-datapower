
data "datapower_gatewaypeeringgroup" "test" {
  depends_on = [datapower_gatewaypeeringgroup.test]
  app_domain = "acc_test_domain"
}