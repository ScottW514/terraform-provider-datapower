
data "datapower_assemblyactionsetvar" "test" {
  depends_on = [datapower_assemblyactionsetvar.test]
  app_domain = "acc_test_domain"
}