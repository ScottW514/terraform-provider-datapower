
resource "datapower_assemblyactionratelimit" "test" {
  id         = "ResTestAssemblyActionRateLimit"
  app_domain = "acceptance_test"
  source     = "plan-default"
}