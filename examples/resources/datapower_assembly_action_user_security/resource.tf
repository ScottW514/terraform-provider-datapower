
resource "datapower_assembly_action_user_security" "test" {
  id                      = "ResTestAssemblyActionUserSecurity"
  app_domain              = "acceptance_test"
  factor_id               = "default"
  extract_identity_method = "basic"
  user_auth_method        = "user-registry"
  user_registry           = "AccTest_APIAuthURLRegistry"
  user_az_method          = "authenticated"
  az_table_default_entry = [{
    name = "tableentryname"
  }]
}