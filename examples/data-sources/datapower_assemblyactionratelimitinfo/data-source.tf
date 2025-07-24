
data "datapower_assemblyactionratelimitinfo" "test" {
  depends_on = [datapower_assemblyactionratelimitinfo.test]
  app_domain = "acc_test_domain"
}