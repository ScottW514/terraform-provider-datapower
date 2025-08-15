
resource "datapower_assembly_action_invoke" "test" {
  id           = "ResTestAssemblyActionInvoke"
  app_domain   = "acceptance_test"
  url          = "https://localhost"
  method       = "Keep"
  backend_type = "detect"
  cache_type   = "Protocol"
}