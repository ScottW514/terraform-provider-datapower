
resource "datapower_assemblyactionsetvar" "test" {
  id         = "AssemblyActionSetVar_name"
  app_domain = "acc_test_domain"
  variable = [{
    name = "varname"
  }]
}