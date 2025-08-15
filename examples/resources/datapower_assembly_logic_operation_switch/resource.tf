
resource "datapower_assembly_logic_operation_switch" "test" {
  id         = "ResTestAssemblyLogicOperationSwitch"
  app_domain = "acceptance_test"
  case = [{
    execute      = "default-api-rule"
    operation_id = "operationid"
  }]
}