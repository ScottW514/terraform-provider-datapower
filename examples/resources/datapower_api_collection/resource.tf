
resource "datapower_api_collection" "test" {
  id         = "ResTestAPICollection"
  app_domain = "acceptance_test"
  org_id     = "orgid"
  org_name   = "orgname"
  routing_prefix = [{
    type = "host"
  }]
  default_rate_limit = [{
    name = "apiratelimit"
    rate = 1000
  }]
  assembly_burst_limit = [{
    name  = "apiburstlimit"
    burst = 1000
  }]
  assembly_rate_limit = [{
    name = "apiratelimit"
    rate = 1000
  }]
  assembly_count_limit = [{
    name  = "countlimit"
    count = 1000
  }]
  api_processing_rule = "default-api-rule"
  api_error_rule      = "default-api-error-rule"
  plan                = ["AccTest_APIPlan"]
  parse_settings_reference = {
    url = "some_url"
  }
}