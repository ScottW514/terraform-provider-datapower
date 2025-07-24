
data "datapower_assemblyactionextract" "test" {
  depends_on = [datapower_assemblyactionextract.test]
  app_domain = "acc_test_domain"
}