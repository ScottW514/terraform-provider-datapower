
resource "datapower_assembly_action_function_call" "test" {
  id            = "ResTestAssemblyActionFunctionCall"
  app_domain    = "acceptance_test"
  function_call = "default-func-global"
}