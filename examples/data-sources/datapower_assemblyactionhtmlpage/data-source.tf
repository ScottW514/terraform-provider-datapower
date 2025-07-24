
data "datapower_assemblyactionhtmlpage" "test" {
  depends_on = [datapower_assemblyactionhtmlpage.test]
  app_domain = "acc_test_domain"
}