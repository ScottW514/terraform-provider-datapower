
resource "datapower_nfsstaticmount" "test" {
  id         = "ResTestNFSStaticMount"
  app_domain = "acceptance_test"
  remote     = "url://test"
}