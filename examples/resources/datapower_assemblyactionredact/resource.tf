
resource "datapower_assemblyactionredact" "test" {
  id         = "ResTestAssemblyActionRedact"
  app_domain = "acceptance_test"
  redact = [{
    path = "path"
  }]
}