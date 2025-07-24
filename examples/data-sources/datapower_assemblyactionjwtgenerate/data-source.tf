
data "datapower_assemblyactionjwtgenerate" "test" {
  depends_on = [datapower_assemblyactionjwtgenerate.test]
  app_domain = "acc_test_domain"
}