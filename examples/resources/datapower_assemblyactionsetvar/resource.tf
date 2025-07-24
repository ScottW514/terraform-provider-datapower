
resource "datapower_assemblyactionsetvar" "test" {
  id         = "_name"
  app_domain = "acc_test_domain"
  variable = [{
    name = "varname"
  }]
}