
data "datapower_durationmonitor" "test" {
  depends_on = [datapower_durationmonitor.test]
  app_domain = "acc_test_domain"
}