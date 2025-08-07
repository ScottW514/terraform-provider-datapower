
resource "datapower_apisecurityrequirement" "test" {
  id         = "ResTestAPISecurityRequirement"
  app_domain = "acceptance_test"
  security   = ["AccTest_APISecurityAPIKey"]
}