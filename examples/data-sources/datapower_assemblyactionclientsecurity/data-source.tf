
data "datapower_assemblyactionclientsecurity" "test" {
  depends_on = [datapower_assemblyactionclientsecurity.test]
  app_domain = "acc_test_domain"
}