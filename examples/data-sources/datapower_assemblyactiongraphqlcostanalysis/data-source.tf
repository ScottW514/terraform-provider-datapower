
data "datapower_assemblyactiongraphqlcostanalysis" "test" {
  depends_on = [datapower_assemblyactiongraphqlcostanalysis.test]
  app_domain = "acc_test_domain"
}