
resource "datapower_b2bxpathroutingpolicy" "test" {
  id              = "B2BXPathRoutingPolicy_name"
  app_domain      = "acc_test_domain"
  sender_x_path   = "senderpath"
  receiver_x_path = "senderpath"
}