
resource "datapower_web_app_response" "test" {
  id          = "ResTest_WebAppResponse"
  app_domain  = "acceptance_test"
  policy_type = "admission"
}