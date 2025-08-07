
resource "datapower_assemblyactiongatewayscript" "test" {
  id         = "ResTestAssemblyActionGatewayScript"
  app_domain = "acceptance_test"
  source     = "gsfile"
}