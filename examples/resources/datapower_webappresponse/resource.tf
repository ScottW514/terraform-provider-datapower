
resource "datapower_webappresponse" "test" {
  id          = "ResTest_WebAppResponse"
  app_domain  = "acceptance_test"
  policy_type = "admission"
}