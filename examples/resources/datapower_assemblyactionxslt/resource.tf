
resource "datapower_assemblyactionxslt" "test" {
  id         = "ResTestAssemblyActionXSLT"
  app_domain = "acceptance_test"
  stylesheet = "local:///stylesheet"
}