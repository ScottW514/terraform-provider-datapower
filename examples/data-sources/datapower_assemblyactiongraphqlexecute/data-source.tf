
data "datapower_assemblyactiongraphqlexecute" "test" {
  depends_on = [datapower_assemblyactiongraphqlexecute.test]
  app_domain = "acc_test_domain"
}