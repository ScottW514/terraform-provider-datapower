
resource "datapower_xpath_routing_map" "test" {
  id         = "ResTestXPathRoutingMap"
  app_domain = "acceptance_test"
  xpath_routing_rules = [{
    xpath = "*"
    host  = "localhost"
    port  = 8888
    ssl   = false
  }]
  name_space_mappings = null
}