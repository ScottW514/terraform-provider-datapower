
resource "datapower_assemblyactionfunctioncall" "test" {
  id            = "ResTestAssemblyActionFunctionCall"
  app_domain    = "acceptance_test"
  function_call = "default-func-global"
}