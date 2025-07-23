
data "datapower_sqldatasource" "test" {
  depends_on = [datapower_sqldatasource.test]
  app_domain = "acc_test_domain"
}