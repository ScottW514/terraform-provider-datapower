
resource "datapower_assembly_action_client_security" "test" {
  id                         = "ResTestAssemblyActionClientSecurity"
  app_domain                 = "acceptance_test"
  id_name                    = "idname"
  secret_name                = "secretname"
  authenticate_client_method = "native"
}