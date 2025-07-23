
resource "datapower_ratelimitdefinition" "test" {
  id         = "RateLimitDefinition_test"
  app_domain = "acc_test_domain"
  type       = "rate"
  rate       = 1000
}