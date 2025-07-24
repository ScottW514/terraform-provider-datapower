
data "datapower_assemblyactionjwtvalidate" "test" {
  depends_on = [datapower_assemblyactionjwtvalidate.test]
  app_domain = "acc_test_domain"
}