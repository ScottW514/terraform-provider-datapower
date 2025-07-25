
data "datapower_lunapartition" "test" {
  depends_on = [datapower_lunapartition.test]
  app_domain = "acc_test_domain"
}