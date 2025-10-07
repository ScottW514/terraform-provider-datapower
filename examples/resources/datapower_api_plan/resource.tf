
resource "datapower_api_plan" "test" {
  id         = "ResTestAPIPlan"
  app_domain = "acceptance_test"
  rate_limit = [{
    name = "apiratelimit"
    rate = 1000
  }]
  burst_limit = [{
    name  = "apiburstlimit"
    burst = 1000
  }]
  assembly_burst_limit = [{
    name  = "apiburstlimit"
    burst = 1000
  }]
  assembly_burst_limit_definition = [{
    short_name = "shortname"
    definition = "AccTest_RateLimitDefinition"
  }]
  assembly_rate_limit = [{
    name = "apiratelimit"
    rate = 1000
  }]
  assembly_rate_limit_definition = [{
    short_name = "shortname"
    definition = "AccTest_RateLimitDefinition"
  }]
  assembly_count_limit = [{
    name  = "countlimit"
    count = 1000
  }]
  assembly_count_limit_definition = [{
    short_name = "shortname"
    definition = "AccTest_RateLimitDefinition"
  }]
  api              = ["AccTest_APIDefinition"]
  rate_limit_scope = "per-application"
}