
data "datapower_assemblyactionvalidate" "test" {
  depends_on = [datapower_assemblyactionvalidate.test]
  app_domain = "acc_test_domain"
}