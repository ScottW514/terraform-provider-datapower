
resource "datapower_xmlfirewallservice" "test" {
  id            = "ResTestXMLFirewallService"
  app_domain    = "acceptance_test"
  type          = "dynamic-backend"
  xml_manager   = "default"
  local_port    = 7575
  local_address = "0.0.0.0"
}