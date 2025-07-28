
resource "datapower_assemblyactionredact" "test" {
  id         = "AssemblyActionRedact_name"
  app_domain = "acc_test_domain"
  redact = [{
    path = "path"
  }]
}