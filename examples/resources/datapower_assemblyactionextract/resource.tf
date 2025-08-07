
resource "datapower_assemblyactionextract" "test" {
  id         = "ResTestAssemblyActionExtract"
  app_domain = "acceptance_test"
  extract = [{
    capture = "capture"
  }]
}