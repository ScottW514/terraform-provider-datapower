
resource "datapower_b2b_xpath_routing_policy" "test" {
  id             = "ResTestB2BXPathRoutingPolicy"
  app_domain     = "acceptance_test"
  sender_xpath   = "senderpath"
  receiver_xpath = "senderpath"
}