
resource "datapower_api_security_requirement" "test" {
  id         = "ResTestAPISecurityRequirement"
  app_domain = "acceptance_test"
  security   = ["AccTest_APISecurityAPIKey"]
}