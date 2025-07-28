
resource "datapower_assemblylogicoperationswitch" "test" {
  id         = "AssemblyLogicOperationSwitch_name"
  app_domain = "acc_test_domain"
  case = [{
    execute = "default-api-rule"
  }]
}