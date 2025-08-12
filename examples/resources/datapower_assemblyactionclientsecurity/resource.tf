
resource "datapower_assemblyactionclientsecurity" "test" {
  id                         = "ResTestAssemblyActionClientSecurity"
  app_domain                 = "acceptance_test"
  id_name                    = "idname"
  secret_name                = "secretname"
  authenticate_client_method = "native"
}