
data "datapower_distributedvariable" "test" {
  depends_on = [datapower_distributedvariable.test]
  app_domain = "acc_test_domain"
}