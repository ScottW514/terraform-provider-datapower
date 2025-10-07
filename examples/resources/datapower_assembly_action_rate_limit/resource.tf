
resource "datapower_assembly_action_rate_limit" "test" {
  id         = "ResTestAssemblyActionRateLimit"
  app_domain = "acceptance_test"
  source     = "plan-default"
  rate_limit = [{
    name = "ratelimitinfoname"
  }]
  count_limit = [{
    name   = "countlimitinfoname"
    action = "inc"
  }]
  rate_limit_definition = [{
    name = "AccTest_RateLimitDefinition"
  }]
}