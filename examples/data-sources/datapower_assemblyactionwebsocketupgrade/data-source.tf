
data "datapower_assemblyactionwebsocketupgrade" "test" {
  depends_on = [datapower_assemblyactionwebsocketupgrade.test]
  app_domain = "acc_test_domain"
}