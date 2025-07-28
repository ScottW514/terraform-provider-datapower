
resource "datapower_assemblylogicswitch" "test" {
  id         = "AssemblyLogicSwitch_name"
  app_domain = "acc_test_domain"
  case = [{
    condition = "condition"
    execute   = "default-api-rule"
  }]
}