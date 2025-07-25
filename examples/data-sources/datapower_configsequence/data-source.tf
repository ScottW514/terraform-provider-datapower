
data "datapower_configsequence" "test" {
  depends_on = [datapower_configsequence.test]
  app_domain = "acc_test_domain"
}