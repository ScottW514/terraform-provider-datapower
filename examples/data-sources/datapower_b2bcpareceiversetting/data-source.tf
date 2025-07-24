
data "datapower_b2bcpareceiversetting" "test" {
  depends_on = [datapower_b2bcpareceiversetting.test]
  app_domain = "acc_test_domain"
}