
data "datapower_schemaexceptionmap" "test" {
  depends_on = [datapower_schemaexceptionmap.test]
  app_domain = "acc_test_domain"
}