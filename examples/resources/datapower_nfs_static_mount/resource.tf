
resource "datapower_nfs_static_mount" "test" {
  id         = "ResTestNFSStaticMount"
  app_domain = "acceptance_test"
  remote     = "url://test"
}