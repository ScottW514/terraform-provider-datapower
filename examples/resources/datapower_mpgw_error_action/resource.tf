
resource "datapower_mpgw_error_action" "test" {
  id         = "ResTestMPGWErrorAction"
  app_domain = "acceptance_test"
  type       = "static"
  local_url  = "store:///schemas/XMLSchema.dtd"
}