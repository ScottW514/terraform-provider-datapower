
data "datapower_assemblyactiongatewayscript" "test" {
  depends_on = [datapower_assemblyactiongatewayscript.test]
  app_domain = "acc_test_domain"
}