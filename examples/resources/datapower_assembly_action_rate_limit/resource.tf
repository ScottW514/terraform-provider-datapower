
resource "datapower_assembly_action_rate_limit" "test" {
  id         = "ResTestAssemblyActionRateLimit"
  app_domain = "acceptance_test"
  source     = "plan-default"
}