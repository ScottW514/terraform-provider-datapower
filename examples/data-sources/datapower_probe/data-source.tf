
data "datapower_probe" "test" {
  depends_on = [datapower_probe.test]
  app_domain = "acc_test_domain"
}