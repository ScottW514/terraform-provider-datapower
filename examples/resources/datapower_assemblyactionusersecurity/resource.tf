
resource "datapower_assemblyactionusersecurity" "test" {
  id                      = "ResTestAssemblyActionUserSecurity"
  app_domain              = "acceptance_test"
  factor_id               = "default"
  extract_identity_method = "basic"
  user_auth_method        = "user-registry"
  user_registry           = "AccTest_APIAuthURLRegistry"
  user_az_method          = "authenticated"
}