
resource "datapower_operationratelimit" "test" {
  id         = "ResTestOperationRateLimit"
  app_domain = "acceptance_test"
  operation  = "AccTest_APIOperation"
}