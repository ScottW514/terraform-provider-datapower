
resource "datapower_slm_policy" "test" {
  id         = "ResTestSLMPolicy"
  app_domain = "acceptance_test"
  statement = [{
    slm_id                                          = 514
    action                                          = "notify"
    maximum_total_reporting_records                 = 5000
    maximum_resources_and_credentials_for_threshold = 5000
  }]
}