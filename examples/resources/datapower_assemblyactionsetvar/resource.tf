
resource "datapower_assemblyactionsetvar" "test" {
  id         = "ResTestAssemblyActionSetVar"
  app_domain = "acceptance_test"
  variable = [{
    name  = "varname"
    value = "value"
  }]
}