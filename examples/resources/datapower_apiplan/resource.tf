
resource "datapower_apiplan" "test" {
  id               = "ResTestAPIPlan"
  app_domain       = "acceptance_test"
  api              = ["AccTest_APIDefinition"]
  rate_limit_scope = "per-application"
}