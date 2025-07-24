
data "datapower_assemblyactioninvoke" "test" {
  depends_on = [datapower_assemblyactioninvoke.test]
  app_domain = "acc_test_domain"
}