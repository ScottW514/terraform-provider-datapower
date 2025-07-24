
data "datapower_assemblyfunction" "test" {
  depends_on = [datapower_assemblyfunction.test]
  app_domain = "acc_test_domain"
}