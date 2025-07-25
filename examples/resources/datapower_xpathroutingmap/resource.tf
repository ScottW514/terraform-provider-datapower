
resource "datapower_xpathroutingmap" "test" {
  id         = "XPathRoutingMap_name"
  app_domain = "acc_test_domain"
  x_path_routing_rules = [{
    x_path = "*"
    host   = "localhost"
    port   = 8888
    ssl    = false
  }]
}