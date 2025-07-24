
data "datapower_assemblyactiongraphqlintrospect" "test" {
  depends_on = [datapower_assemblyactiongraphqlintrospect.test]
  app_domain = "acc_test_domain"
}