
resource "datapower_assembly_function" "test" {
  id         = "ResTestAssemblyFunction"
  app_domain = "acceptance_test"
  parameter = [{
    name = "assemblyfunctionparametername"
  }]
}