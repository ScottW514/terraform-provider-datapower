
data "datapower_mpgwerroraction" "test" {
  depends_on = [datapower_mpgwerroraction.test]
  app_domain = "acc_test_domain"
}