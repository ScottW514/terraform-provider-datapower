
data "datapower_apifinal" "test" {
  depends_on = [datapower_apifinal.test]
  app_domain = "acc_test_domain"
}