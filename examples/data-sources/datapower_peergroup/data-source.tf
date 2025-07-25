
data "datapower_peergroup" "test" {
  depends_on = [datapower_peergroup.test]
  app_domain = "acc_test_domain"
}