
data "datapower_assemblyactionxslt" "test" {
  depends_on = [datapower_assemblyactionxslt.test]
  app_domain = "acc_test_domain"
}