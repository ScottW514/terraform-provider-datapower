
resource "datapower_b2bxpathroutingpolicy" "test" {
  id              = "ResTestB2BXPathRoutingPolicy"
  app_domain      = "acceptance_test"
  sender_x_path   = "senderpath"
  receiver_x_path = "senderpath"
}