
data "datapower_assemblyactionredact" "test" {
  depends_on = [datapower_assemblyactionredact.test]
  app_domain = "acc_test_domain"
}