
resource "datapower_assembly_action_gateway_script" "test" {
  id         = "ResTestAssemblyActionGatewayScript"
  app_domain = "acceptance_test"
  source     = "gsfile"
}