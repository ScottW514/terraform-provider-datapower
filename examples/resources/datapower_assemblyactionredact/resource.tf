
resource "datapower_assemblyactionredact" "test" {
  id         = "_name"
  app_domain = "acc_test_domain"
  redact = [{
    path = "path"
  }]
}