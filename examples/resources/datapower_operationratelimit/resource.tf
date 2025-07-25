
resource "datapower_operationratelimit" "test" {
  id         = "OperationRateLimit_name"
  app_domain = "acc_test_domain"
  operation  = datapower_apioperation.test.id
}