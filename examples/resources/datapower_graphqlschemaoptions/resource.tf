
resource "datapower_graphqlschemaoptions" "test" {
  id         = "GraphQLSchemaOptions_name"
  app_domain = "acc_test_domain"
  api        = datapower_apidefinition.test.id
}