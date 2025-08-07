
resource "datapower_corspolicy" "test" {
  id         = "ResTestCORSPolicy"
  app_domain = "acceptance_test"
  rule       = ["AccTest_CORSRule"]
}