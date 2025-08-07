
resource "datapower_peergroup" "test" {
  id         = "ResTestPeerGroup"
  app_domain = "acceptance_test"
  type       = "slm"
}