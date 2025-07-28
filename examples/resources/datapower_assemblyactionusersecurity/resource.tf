
resource "datapower_assemblyactionusersecurity" "test" {
  id                      = "AssemblyActionUserSecurity_name"
  app_domain              = "acc_test_domain"
  factor_id               = "default"
  extract_identity_method = "basic"
  user_auth_method        = "user-registry"
  user_az_method          = "authenticated"
}