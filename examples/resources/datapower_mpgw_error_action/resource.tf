
resource "datapower_mpgw_error_action" "test" {
  id         = "ResTestMPGWErrorAction"
  app_domain = "acceptance_test"
  type       = "static"
  local_url  = "store:///schemas/XMLSchema.dtd"
  header_injection = [{
    header_tag_value = "HEADER_VALUE"
  }]
}