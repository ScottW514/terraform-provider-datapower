
data "datapower_operationratelimit" "test" {
  depends_on = [datapower_operationratelimit.test]
  app_domain = "acc_test_domain"
}