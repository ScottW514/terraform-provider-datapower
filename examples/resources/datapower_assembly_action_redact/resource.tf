
resource "datapower_assembly_action_redact" "test" {
  id         = "ResTestAssemblyActionRedact"
  app_domain = "acceptance_test"
  redact = [{
    path = "path"
  }]
}