
resource "datapower_assembly_action_xslt" "test" {
  id         = "ResTestAssemblyActionXSLT"
  app_domain = "acceptance_test"
  stylesheet = "local:///stylesheet"
}