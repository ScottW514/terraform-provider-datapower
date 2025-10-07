
resource "datapower_assembly_action_parse" "test" {
  id         = "ResTestAssemblyActionParse"
  app_domain = "acceptance_test"
  parse_settings_reference = {
    url = "some_url"
  }
}