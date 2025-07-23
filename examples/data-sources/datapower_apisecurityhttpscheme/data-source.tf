
data "datapower_apisecurityhttpscheme" "test" {
  depends_on = [datapower_apisecurityhttpscheme.test]
  app_domain = "acc_test_domain"
}