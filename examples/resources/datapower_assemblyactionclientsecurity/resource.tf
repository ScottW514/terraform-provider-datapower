
resource "datapower_assemblyactionclientsecurity" "test" {
  id                         = "ResTestAssemblyActionClientSecurity"
  app_domain                 = "acceptance_test"
  authenticate_client_method = "native"
}