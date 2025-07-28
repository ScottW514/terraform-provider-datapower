
resource "datapower_assemblyactiongatewayscript" "test" {
  id         = "AssemblyActionGatewayScript_name"
  app_domain = "acc_test_domain"
  source     = "gsfile"
}