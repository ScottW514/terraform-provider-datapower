
resource "datapower_assemblyactionclientsecurity" "test" {
  id                         = "AssemblyActionClientSecurity_name"
  app_domain                 = "acc_test_domain"
  authenticate_client_method = "native"
}