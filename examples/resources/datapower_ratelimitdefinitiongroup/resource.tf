
resource "datapower_ratelimitdefinitiongroup" "test" {
  id               = "ResTestRateLimitDefinitionGroup"
  app_domain       = "acceptance_test"
  update_on_exceed = "all"
}