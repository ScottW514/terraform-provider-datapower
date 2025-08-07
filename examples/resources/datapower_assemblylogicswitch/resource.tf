
resource "datapower_assemblylogicswitch" "test" {
  id         = "ResTestAssemblyLogicSwitch"
  app_domain = "acceptance_test"
  case = [{
    condition = "condition"
    execute   = "default-api-rule"
  }]
}