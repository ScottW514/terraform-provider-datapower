
data "datapower_ftpquotecommands" "test" {
  depends_on = [datapower_ftpquotecommands.test]
  app_domain = "acc_test_domain"
}