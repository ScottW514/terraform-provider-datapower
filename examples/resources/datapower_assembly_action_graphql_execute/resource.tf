
resource "datapower_assembly_action_graphql_execute" "test" {
  id         = "ResTestAssemblyActionGraphQLExecute"
  app_domain = "acceptance_test"
  target_map_rule = [{
    target  = "target"
    execute = "default-func-call-global"
  }]
}