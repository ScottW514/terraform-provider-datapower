
resource "datapower_assemblyactionlog" "test" {
  id         = "ResTestAssemblyActionLog"
  app_domain = "acceptance_test"
  mode       = "gather-only"
}