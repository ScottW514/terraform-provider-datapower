
resource "datapower_apicollection" "test" {
  id         = "APICollection_test"
  app_domain = "acc_test_domain"
  org_id     = "orgid"
  org_name   = "orgname"
  routing_prefix = [{
    type = "host"
  }]
  api_processing_rule = "default-api-rule"
  api_error_rule      = "default-api-error-rule"
  plan                = [datapower_apiplan.test.id]
}