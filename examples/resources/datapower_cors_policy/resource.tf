
resource "datapower_cors_policy" "test" {
  id         = "ResTestCORSPolicy"
  app_domain = "acceptance_test"
  rule       = ["AccTest_CORSRule"]
}