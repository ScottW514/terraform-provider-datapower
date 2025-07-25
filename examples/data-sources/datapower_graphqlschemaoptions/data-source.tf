
data "datapower_graphqlschemaoptions" "test" {
  depends_on = [datapower_graphqlschemaoptions.test]
  app_domain = "acc_test_domain"
}