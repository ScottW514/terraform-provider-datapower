
resource "datapower_assembly" "test" {
  id         = "ResTestAssembly"
  app_domain = "acceptance_test"
  rule       = "default-empty-rule"
  catch = [{
    error   = "errorname"
    handler = "default-empty-rule"
  }]
}