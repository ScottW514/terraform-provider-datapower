
resource "datapower_peer_group" "test" {
  id         = "ResTestPeerGroup"
  app_domain = "acceptance_test"
  type       = "slm"
  url        = ["http://localhost"]
}