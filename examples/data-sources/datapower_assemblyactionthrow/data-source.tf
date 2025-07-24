
data "datapower_assemblyactionthrow" "test" {
  depends_on = [datapower_assemblyactionthrow.test]
  app_domain = "acc_test_domain"
}