
resource "datapower_operation_rate_limit" "test" {
  id         = "ResTestOperationRateLimit"
  app_domain = "acceptance_test"
  operation  = "AccTest_APIOperation"
}