
data "datapower_countmonitor" "test" {
  depends_on = [datapower_countmonitor.test]
  app_domain = "acc_test_domain"
}