
resource "datapower_rate_limit_configuration" "test" {
  app_domain = "acceptance_test"
  parameters = [{
    name  = "ratelimitname"
    value = "ratelimitvalue"
  }]
}