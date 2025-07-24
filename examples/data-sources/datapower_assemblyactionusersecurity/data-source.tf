
data "datapower_assemblyactionusersecurity" "test" {
  depends_on = [datapower_assemblyactionusersecurity.test]
  app_domain = "acc_test_domain"
}