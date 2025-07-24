
data "datapower_assemblyactionmap" "test" {
  depends_on = [datapower_assemblyactionmap.test]
  app_domain = "acc_test_domain"
}