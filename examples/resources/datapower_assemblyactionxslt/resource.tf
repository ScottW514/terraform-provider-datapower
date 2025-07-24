
resource "datapower_assemblyactionxslt" "test" {
  id         = "_name"
  app_domain = "acc_test_domain"
  stylesheet = "local:///stylesheet"
}