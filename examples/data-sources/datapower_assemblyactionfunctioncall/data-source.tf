
data "datapower_assemblyactionfunctioncall" "test" {
  depends_on = [datapower_assemblyactionfunctioncall.test]
  app_domain = "acc_test_domain"
}