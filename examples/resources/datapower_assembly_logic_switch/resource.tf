
resource "datapower_assembly_logic_switch" "test" {
  id         = "ResTestAssemblyLogicSwitch"
  app_domain = "acceptance_test"
  case = [{
    condition = "condition"
    execute   = "default-api-rule"
  }]
}