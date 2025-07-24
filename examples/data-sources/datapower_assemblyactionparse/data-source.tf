
data "datapower_assemblyactionparse" "test" {
  depends_on = [datapower_assemblyactionparse.test]
  app_domain = "acc_test_domain"
}