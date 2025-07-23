
data "datapower_apirule" "test" {
  depends_on = [datapower_apirule.test]
  app_domain = "acc_test_domain"
}