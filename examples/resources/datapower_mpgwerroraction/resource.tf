
resource "datapower_mpgwerroraction" "test" {
  id         = "ResTestMPGWErrorAction"
  app_domain = "acceptance_test"
  type       = "static"
  local_url  = "store:///schemas/XMLSchema.dtd"
}