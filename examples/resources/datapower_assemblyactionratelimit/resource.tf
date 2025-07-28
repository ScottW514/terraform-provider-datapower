
resource "datapower_assemblyactionratelimit" "test" {
  id         = "AssemblyActionRateLimit_name"
  app_domain = "acc_test_domain"
  source     = "plan-default"
}