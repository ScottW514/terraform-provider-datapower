
resource "datapower_rate_limit_definition" "test" {
  id         = "ResTestRateLimitDefinition"
  app_domain = "acceptance_test"
  type       = "rate"
  rate       = 1000
  parameters = [{
    name  = "name"
    value = "value"
  }]
}