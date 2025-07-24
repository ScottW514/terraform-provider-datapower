
data "datapower_assembly" "test" {
  depends_on = [datapower_assembly.test]
  app_domain = "acc_test_domain"
}