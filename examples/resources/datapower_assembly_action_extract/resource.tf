
resource "datapower_assembly_action_extract" "test" {
  id         = "ResTestAssemblyActionExtract"
  app_domain = "acceptance_test"
  extract = [{
    capture = "capture"
  }]
}