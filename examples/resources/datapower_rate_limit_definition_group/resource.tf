
resource "datapower_rate_limit_definition_group" "test" {
  id               = "ResTestRateLimitDefinitionGroup"
  app_domain       = "acceptance_test"
  update_on_exceed = "all"
}