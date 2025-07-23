
data "datapower_statistics" "test" {
  depends_on = [datapower_statistics.test]
  app_domain = "acc_test_domain"
}