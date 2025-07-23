
resource "datapower_xmlfirewallservice" "test" {
  id            = "XMLFirewallService_name"
  app_domain    = "acc_test_domain"
  type          = "dynamic-backend"
  xml_manager   = "default"
  local_port    = 8888
  local_address = "0.0.0.0"
}