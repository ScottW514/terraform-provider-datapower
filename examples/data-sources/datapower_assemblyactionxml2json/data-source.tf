
data "datapower_assemblyactionxml2json" "test" {
  depends_on = [datapower_assemblyactionxml2json.test]
  app_domain = "acc_test_domain"
}