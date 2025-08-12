
resource "datapower_assemblylogicoperationswitch" "test" {
  id         = "ResTestAssemblyLogicOperationSwitch"
  app_domain = "acceptance_test"
  case = [{
    execute      = "default-api-rule"
    operation_id = "operationid"
  }]
}