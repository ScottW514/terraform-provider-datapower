
data "datapower_loglabel" "test" {
  depends_on = [datapower_loglabel.test]
  app_domain = "acc_test_domain"
}