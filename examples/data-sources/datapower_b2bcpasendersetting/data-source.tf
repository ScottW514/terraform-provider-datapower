
data "datapower_b2bcpasendersetting" "test" {
  depends_on = [datapower_b2bcpasendersetting.test]
  app_domain = "acc_test_domain"
}