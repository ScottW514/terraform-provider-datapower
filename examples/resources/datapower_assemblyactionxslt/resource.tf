
resource "datapower_assemblyactionxslt" "test" {
  id         = "AssemblyActionXSLT_name"
  app_domain = "acc_test_domain"
  stylesheet = "local:///stylesheet"
}