
resource "datapower_apisecuritybasicauth" "test" {
  id            = "ResTestAPISecurityBasicAuth"
  app_domain    = "acceptance_test"
  user_registry = "AccTest_APIAuthURLRegistry"
}