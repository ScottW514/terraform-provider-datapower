
data "datapower_assemblylogicoperationswitch" "test" {
  depends_on = [datapower_assemblylogicoperationswitch.test]
  app_domain = "acc_test_domain"
}