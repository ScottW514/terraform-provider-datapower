
data "datapower_nfsstaticmount" "test" {
  depends_on = [datapower_nfsstaticmount.test]
  app_domain = "acc_test_domain"
}