
resource "datapower_xpath_routing_map" "test" {
  id         = "ResTestXPathRoutingMap"
  app_domain = "acceptance_test"
  x_path_routing_rules = [{
    x_path = "*"
    host   = "localhost"
    port   = 8888
    ssl    = false
  }]
}