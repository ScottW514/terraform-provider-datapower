
data "datapower_xpathroutingmap" "test" {
  depends_on = [datapower_xpathroutingmap.test]
  app_domain = "acc_test_domain"
}