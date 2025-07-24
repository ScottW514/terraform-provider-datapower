
data "datapower_mqmanagergroup" "test" {
  depends_on = [datapower_mqmanagergroup.test]
  app_domain = "acc_test_domain"
}