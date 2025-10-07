
resource "datapower_assembly_action_map" "test" {
  id         = "ResTestAssemblyActionMap"
  app_domain = "acceptance_test"
  location   = "local:///file"
  parse_settings_reference = {
    url = "some_url"
  }
}