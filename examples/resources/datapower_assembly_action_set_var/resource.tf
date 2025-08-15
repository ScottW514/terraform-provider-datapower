
resource "datapower_assembly_action_set_var" "test" {
  id         = "ResTestAssemblyActionSetVar"
  app_domain = "acceptance_test"
  variable = [{
    name  = "varname"
    value = "value"
  }]
}