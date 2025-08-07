
resource "datapower_apicollection" "test" {
  id         = "ResTestAPICollection"
  app_domain = "acceptance_test"
  org_id     = "orgid"
  org_name   = "orgname"
  routing_prefix = [{
    type = "host"
  }]
  api_processing_rule = "default-api-rule"
  api_error_rule      = "default-api-error-rule"
  plan                = ["AccTest_APIPlan"]
}