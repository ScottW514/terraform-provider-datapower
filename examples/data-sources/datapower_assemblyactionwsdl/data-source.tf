
data "datapower_assemblyactionwsdl" "test" {
  depends_on = [datapower_assemblyactionwsdl.test]
  app_domain = "acc_test_domain"
}