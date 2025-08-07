
resource "datapower_ratelimitdefinition" "test" {
  id         = "ResTestRateLimitDefinition"
  app_domain = "acceptance_test"
  type       = "rate"
  rate       = 1000
}