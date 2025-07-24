
resource "datapower_assemblyactionfunctioncall" "test" {
  id            = "_name"
  app_domain    = "acc_test_domain"
  function_call = datapower_assemblyfunction.test.id
}