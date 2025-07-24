
data "datapower_assemblyactionoauth" "test" {
  depends_on = [datapower_assemblyactionoauth.test]
  app_domain = "acc_test_domain"
}