
resource "datapower_api_security_basic_auth" "test" {
  id            = "ResTestAPISecurityBasicAuth"
  app_domain    = "acceptance_test"
  user_registry = "AccTest_APIAuthURLRegistry"
}