
data "datapower_assemblyactionlog" "test" {
  depends_on = [datapower_assemblyactionlog.test]
  app_domain = "acc_test_domain"
}