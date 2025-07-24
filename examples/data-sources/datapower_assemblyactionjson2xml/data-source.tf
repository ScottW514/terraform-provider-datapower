
data "datapower_assemblyactionjson2xml" "test" {
  depends_on = [datapower_assemblyactionjson2xml.test]
  app_domain = "acc_test_domain"
}