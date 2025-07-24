
data "datapower_assemblylogicswitch" "test" {
  depends_on = [datapower_assemblylogicswitch.test]
  app_domain = "acc_test_domain"
}