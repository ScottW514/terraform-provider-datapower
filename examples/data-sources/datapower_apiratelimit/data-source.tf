
data "datapower_apiratelimit" "test" {
  depends_on = [datapower_apiratelimit.test]
  app_domain = "acc_test_domain"
}