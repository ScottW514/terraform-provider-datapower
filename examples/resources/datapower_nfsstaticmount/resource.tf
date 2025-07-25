
resource "datapower_nfsstaticmount" "test" {
  id         = "NFSStaticMount_name"
  app_domain = "acc_test_domain"
  remote     = "url://test"
}